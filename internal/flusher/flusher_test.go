package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/timestamp"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozoncp/ocp-timeline-api/internal/flusher"
	"github.com/ozoncp/ocp-timeline-api/internal/mocks"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		crtl     *gomock.Controller
		mockRepo *mocks.MockRepo
		ctx      context.Context
	)

	BeforeEach(func() {
		crtl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(crtl)
		ctx = context.TODO()
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

			It("without a remainder AddEntities must called 4 times", func() {
				mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Times(4)

				flusher := flusher.NewFlusher(2, mockRepo)

				result := flusher.Flush(ctx, data[:4])

				gomega.Expect(result).Should(gomega.BeEquivalentTo(data[:4]))
			})
		})

		Context("Empty input", func() {
			It("input length equal 0  AddEntities must called 0 times", func() {

				input := make([]models.Timeline, 0)

				mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Times(0)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(ctx, input)
			})

			It("input is nil AddEntities must called 0 times", func() {

				var input []models.Timeline

				mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Times(0)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(ctx, input)
			})
		})
	})

	Describe("error processing", func() {
		Context("error processing from Repo", func() {
			It("Repo AddEntities not return error, correct work flusher", func() {
				mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Times(1)

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(ctx, []models.Timeline{newTimeline(1)})
			})

			It("Repo AddEntities return error, flusher throw panic", func() {
				defer func() {
					if rec := recover(); rec == nil {
						Fail("flusher must throw panic")
					}
				}()

				mockRepo.EXPECT().
					AddEntities(ctx, gomock.Any()).
					Return(errors.New("some error"))

				flusher := flusher.NewFlusher(2, mockRepo)

				flusher.Flush(ctx, []models.Timeline{newTimeline(1)})
			})
		})
	})
})

func newTimeline(id uint64) models.Timeline {
	return models.Timeline{
		Id:     id,
		UserId: uint64(2),
		Type:   uint64(3),
		From:   timestamp.Timestamp{Seconds: 100},
		To:     timestamp.Timestamp{Seconds: 200},
	}
}
