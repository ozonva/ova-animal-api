//+build integration

package api_test

// Этот файл скопирован с моего pet-project-а (https://github.com/and-hom/wwmap), потому что писать
// такое второй раз нет никакого смысла
// Возможно, поднимать контейнер именно так не совсем рационально, но ничего лучше я уже не успеваю сделать

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/and-hom/godbt"
	"github.com/and-hom/godbt/contract"
	"github.com/ory/dockertest"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"ova-animal-api/internal/config"
	"path/filepath"
	"strconv"
	"text/template"
)

type DaoTester struct {
	Db       *sql.DB
	Tester   *godbt.Tester
	Resource *dockertest.Resource
	Pool     *dockertest.Pool
}

func (this *DaoTester) Init(db *config.Db) {
	var err error
	this.Pool, err = dockertest.NewPool("")
	if err != nil {
		log.Panic().Msgf("Could not connect to docker: %s", err)
	}
	log.Info().Msg("Connected to docker")

	// pulls an image, creates a container based on it and runs it
	this.Resource, err = this.Pool.Run("postgres", "latest", []string{
		"POSTGRES_DB=" + db.Name,
		"POSTGRES_USER=" + db.Login,
		"POSTGRES_PASSWORD=" + db.Password,
	})
	if err != nil {
		log.Panic().Msgf("Could not start resource: %s", err)
	}
	port, err := strconv.Atoi(this.Resource.GetPort("5432/tcp"))
	if err != nil {
		log.Panic().Msgf("Could not parse port %s: %s", this.Resource.GetPort("5432/tcp"), err)
	}
	log.Info().Msgf("Postgres started on port %d", port)
	db.Port = port

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.Login,
		db.Password,
		db.Host,
		port,
		db.Name,
	)

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := this.Pool.Retry(func() error {
		var err error
		this.Db, err = sql.Open("postgres", connString)
		if err != nil {
			return err
		}
		return this.Db.Ping()
	}); err != nil {
		log.Panic().Msgf("Could not connect to docker: %s", err)
	}
	log.Info().Msg("Connected to database")

	absMigrationsPath, err := filepath.Abs("../../migrations")
	if err != nil {
		log.Panic().Msgf("Could get migrations path: %s", err)
	}
	log.Info().Msgf("Loading migrations from %s", absMigrationsPath)
	err = goose.Up(
		this.Db,
		absMigrationsPath,
	)
	if err != nil {
		log.Panic().Msgf("Could apply migrations: %s", err)
	}
	log.Info().Msg("Migrations applied")

	this.Tester, err = godbt.GetTester(contract.InstallerConfig{
		Type:        "postgres",
		ConnString:  connString,
		ClearMethod: contract.ClearMethodDeleteAll,
	})
	if err != nil {
		log.Panic().Msgf("Could init dbunit: %s", err)
	}
	log.Info().Msg("DBUnit initialized")
}

func (this *DaoTester) Close() error {
	return this.Pool.Purge(this.Resource)
}

func (this *DaoTester) ClearTable(table string) {
	_, err := this.Db.Exec("DELETE FROM " + table)
	if err != nil {
		log.Panic().Msgf("Can't clear table %s: %s", table, err)
	}
}

func (this *DaoTester) ApplyDbunitData(path string) {
	image := this.loadImg(path)
	err := this.Tester.GetInstaller().InstallImage(image)
	if err != nil {
		log.Panic().Msgf("Can't apply dbunit data from %s: %s", path, err)
	}
}

func (this *DaoTester) TestDatabase(table string, path string, opts ...interface{}) {
	var params map[string]string
	if len(opts) > 0 {
		ok := false
		params, ok = opts[0].(map[string]string)
		if !ok {
			log.Panic().Msgf("Param should be map[string]string but was %v", opts[0])
		}
	} else {
		params = make(map[string]string)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Panic().Msgf("Could get path: %s", err)
	}

	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Panic().Msgf("Can't load file %s: %s", path, err)
	}

	tmpl, err := template.New("replace").Parse(string(content))
	if err != nil {
		log.Panic().Msgf("Can't load template: %s", err)
	}
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Panic().Msgf("Can't apply template: %s", err)
	}

	expectedImage, err := this.Tester.GetImageManager().LoadImage(buf.String())
	if err != nil {
		log.Panic().Msgf("Can't load dbunit data from xml %s: %s", buf.String(), err)
	}
	actualImage, err := this.Tester.GetInstaller().GetTableImage(table)
	if err != nil {
		log.Panic().Msgf("Can't load table data from %s: %s", path, err)
	}
	diffs := this.Tester.GetImageManager().GetImagesDiff(expectedImage, actualImage)
	if len(diffs) > 0 {
		log.Panic().Msgf("Tables are different: %v", diffs)
	}
}

func (this *DaoTester) loadImg(path string) contract.Image {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Panic().Msgf("Could get path: %s", err)
	}

	image, err := this.Tester.GetImageManager().LoadImage(absPath)
	if err != nil {
		log.Panic().Msgf("Can't load dbunit data from %s: %s", path, err)
	}
	return image
}
