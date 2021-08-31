package animal_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"ova-animal-api/internal/domain"
	animal_flusher "ova-animal-api/internal/flusher/animal"
	animal_saver "ova-animal-api/internal/saver/animal"
	"time"
)

const capacity = 3

var _ = Describe("Animal", func() {
	var (
		mockCtrl *gomock.Controller
		flusher  *animal_flusher.MockFlusher
		saver    animal_saver.Saver

		animal1 = domain.Animal{1, 1, "Thomas", domain.CAT}
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		flusher = animal_flusher.NewMockFlusher(mockCtrl)
		saver = animal_saver.New(capacity, flusher)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("write", func() {
		It("closed", func() {
			log.Println("Test write in closed saver")
			Expect(saver.Close()).To(BeNil())

			Expect(func() {
				saver.Save(animal1)
			}).To(Panic())
		})

		It("flush on close", func() {
			log.Println("Test flush by closing")
			flusher.
				EXPECT().
				Flush(gomock.Eq([]domain.Animal{animal1})).
				Return([]domain.Animal{})

			// Этот тест может быть хрупким, но мы понадеемся, что время
			// от вызова BeforeEach до Close меньше 1 секунды
			saver.Save(animal1)
			Expect(saver.Close()).To(BeNil())
		})

		It("flush by timeout", func() {
			log.Println("Test flush by timeout")

			flusher.
				EXPECT().
				Flush(gomock.Eq([]domain.Animal{animal1})).
				Return([]domain.Animal{})

			saver.Save(animal1)
			time.Sleep(2 * time.Second)
			Expect(saver.Close()).To(BeNil())
		})
	})

})
