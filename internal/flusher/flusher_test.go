package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-timeline-api/internal/flusher"
	"github.com/ozoncp/ocp-timeline-api/internal/mocks"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		crtl     *gomock.Controller
		mockRepo *mocks.MockRepo
	)

	BeforeEach(func() {
		crtl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(crtl)
	})

	AfterEach(func() {
		crtl.Finish()
	})

	Describe("flusher chunks", func() {
		Context("Not empty input", func() {
			data := []models.Timeline{
				newTimeline(1),
				newTimeline(2),
				newTimeline(3),
				newTimeline(4),
				newTimeline(5),
			}

			It("without a remainder AddEntities must called 2 times", func() {
				mockRepo.EXPECT().AddEntities(gomock.Any()).Times(2)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(data[:4])
			})

			It("with a remainder AddEntities must called 3 times", func() {
				mockRepo.EXPECT().AddEntities(data[:2]).Times(1)
				mockRepo.EXPECT().AddEntities(data[2:4]).Times(1)
				mockRepo.EXPECT().AddEntities(data[4:]).Times(1)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(data)
			})
		})

		Context("Empty input", func() {
			It("input length equal 0  AddEntities must called 0 times", func() {

				input := make([]models.Timeline, 0)

				mockRepo.EXPECT().AddEntities(input).Times(0)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(input)
			})

			It("input is nil AddEntities must called 0 times", func() {

				var input []models.Timeline

				mockRepo.EXPECT().AddEntities(input).Times(0)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(input)
			})
		})
	})

	Describe("error processing", func() {
		Context("error processing from Repo", func() {
			It("Repo AddEntities not return error, correct work flusher", func() {
				mockRepo.EXPECT().AddEntities(gomock.Any()).Times(1)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush([]models.Timeline{newTimeline(1)})
			})

			It("Repo AddEntities return error, flusher throw panic", func() {
				defer func() {
					if rec := recover(); rec == nil {
						Fail("flusher must throw panic")
					}
				}()

				mockRepo.EXPECT().
					AddEntities(gomock.Any()).
					Return(errors.New("some error"))

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush([]models.Timeline{newTimeline(1)})
			})
		})
	})
})

func newTimeline(id uint64) models.Timeline {
	return models.Timeline{
		Id:     id,
		UserId: uint64(2),
		Type:   uint64(3),
		From:   models.Timestamp(int64(200)),
		To:     models.Timestamp(int64(200)),
	}
}
