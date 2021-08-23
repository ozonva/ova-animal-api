package animal_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-animal-api/internal/domain"
	animal_flusher "ova-animal-api/internal/flusher/animal"
	animal_repo "ova-animal-api/internal/repo/animal"
	"ova-animal-api/internal/utils/contains_matcher"
)

var _ = Describe("Animal Flusher", func() {
	var (
		mockCtrl *gomock.Controller
		repo     *animal_repo.MockRepo
		flusher  animal_flusher.Flusher

		animal1 = domain.Animal{1, 1, "Thomas", domain.CAT}
		animal2 = domain.Animal{1, 1, "Jerry", domain.MOUSE}
		animal3 = domain.Animal{1, 1, "Spike", domain.DOG}
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repo = animal_repo.NewMockRepo(mockCtrl)
		flusher = animal_flusher.New(repo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("empty collection", func() {
		emptyAnimals := []domain.Animal{}

		It("Empty", func() {
			repo.
				EXPECT().
				AddEntities(gomock.Eq(emptyAnimals)).
				Return(nil)

			nonFlushed := flusher.Flush(emptyAnimals)
			Expect(nonFlushed).To(BeEmpty())
		})
		It("Nil", func() {
			repo.
				EXPECT().
				AddEntities(gomock.Nil()).
				Return(nil)

			nonFlushed := flusher.Flush(nil)
			Expect(nonFlushed).To(BeEmpty())
		})
	})

	When("single element collection", func() {
		singleAnimal := []domain.Animal{animal1}

		It("Success", func() {
			repo.
				EXPECT().
				AddEntities(gomock.Eq(singleAnimal)).
				Return(nil)

			nonFlushed := flusher.Flush(singleAnimal)
			Expect(nonFlushed).To(BeEmpty())
		})
		It("Fail", func() {
			repo.
				EXPECT().
				AddEntities(gomock.Eq(singleAnimal)).
				Return(errors.New("error"))

			nonFlushed := flusher.Flush(singleAnimal)
			Expect(nonFlushed).To(BeEquivalentTo(singleAnimal))
		})
	})

	When("multiple element collection", func() {
		allAnimals := []domain.Animal{animal1, animal2, animal3}

		It("Success", func() {
			repo.
				EXPECT().
				AddEntities(gomock.Eq(allAnimals)).
				Return(nil)

			nonFlushed := flusher.Flush(allAnimals)
			Expect(nonFlushed).To(BeEmpty())
		})
		It("all failed", func() {
			repo.
				EXPECT().
				AddEntities(gomock.Any()).
				Return(errors.New("error")).
				AnyTimes()

			nonFlushed := flusher.Flush(allAnimals)
			Expect(nonFlushed).To(BeEquivalentTo(allAnimals))
		})
		It("partial failed", func() {
			containsAnimal2 := contains_matcher.New(animal2)

			repo.
				EXPECT().
				AddEntities(containsAnimal2).
				Return(errors.New("error")).
				AnyTimes()
			repo.
				EXPECT().
				AddEntities(gomock.Not(containsAnimal2)).
				Return(nil).
				AnyTimes()

			nonFlushed := flusher.Flush(allAnimals)
			Expect(nonFlushed).To(BeEquivalentTo([]domain.Animal{animal2}))
		})
	})
})
