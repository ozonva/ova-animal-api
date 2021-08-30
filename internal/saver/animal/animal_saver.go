package animal

import (
	"io"
	"log"
	"ova-animal-api/internal/domain"
	"ova-animal-api/internal/flusher/animal"
	"sync"
	"time"
)

const bufCapacity = 10

type Saver interface {
	io.Closer
	Save(entity domain.Animal)
}

func New(capacity uint, flusher animal.Flusher) Saver {
	writeChan := make(chan domain.Animal, 0)
	closeChan := make(chan bool, 0)

	saver := animalSaver{
		flusher:   flusher,
		capacity:  capacity,
		buf:       make([]domain.Animal, 0, bufCapacity),
		ticker:    time.NewTicker(time.Second),
		mutex:     &(sync.Mutex{}),
		writeChan: &writeChan,
		closeChan: &closeChan,
	}
	saver.init()
	return &saver
}

type animalSaver struct {
	flusher  animal.Flusher
	capacity uint

	buf []domain.Animal

	ticker *time.Ticker
	mutex  *sync.Mutex

	writeChan *chan domain.Animal
	closeChan *chan bool
}

func (this *animalSaver) init() {
	go func() {
		for {
			select {
			case <-this.ticker.C:
				log.Println("Tick!")
				this.flush()
			case <-*this.closeChan:
				log.Println("Flush and exit")
				this.flush()
				log.Println("Flushing loop completed")
				return
			case e, ok := <-*this.writeChan:
				if ok {
					this.append(e)
				}
			}
		}

	}()
}

func (this *animalSaver) Save(entity domain.Animal) {
	log.Printf("Schedule entity add %d", entity.Id)
	*this.writeChan <- entity
}

func (this *animalSaver) Close() error {
	log.Println("Closing")
	close(*this.writeChan)
	this.ticker.Stop()
	*this.closeChan <- true
	close(*this.closeChan)
	log.Println("Closed")
	return nil
}

func (this *animalSaver) append(entity domain.Animal) {
	this.withLock(func() {
		log.Printf("Add entity %d. Buf length is %d", entity.Id, len(this.buf))
		this.buf = append(this.buf, entity)
	})
}

func (this *animalSaver) flush() {
	this.withLock(func() {
		if len(this.buf) == 0 {
			log.Println("Skip flushing because of 0 elements")
			return
		}
		log.Printf("Flushing %d elements", len(this.buf))
		prev := this.buf
		result := this.flusher.Flush(prev)
		log.Printf("Flushed. Was %d became %d\n", len(prev), len(result))
		this.buf = result
	})
}

func (this *animalSaver) withLock(payload func()) {
	log.Println("Acquire lock")
	this.mutex.Lock()
	defer this.mutex.Unlock()
	defer log.Println("Release lock")

	payload()
}
