package api_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozoncp/ocp-timeline-api/internal/api"
	"github.com/ozoncp/ocp-timeline-api/internal/mocks"
	"github.com/ozoncp/ocp-timeline-api/internal/models"
	"github.com/ozoncp/ocp-timeline-api/internal/utils"
	desc "github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api"
	"time"
)

var _ = Describe("Api", func() {
	var (
		crtl     *gomock.Controller
		mockRepo *mocks.MockRepo
		ctx      context.Context
		server   desc.OcpTimelineApiServer
	)

	BeforeEach(func() {
		crtl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(crtl)
		ctx = context.TODO()
		server = api.NewServiceOcpTimeline(mockRepo)
	})

	AfterEach(func() {
		crtl.Finish()
	})

	Describe("crud", func() {

		from := utils.ConvertTimeInTimeStamp(time.Now())
		to := utils.ConvertTimeInTimeStamp(time.Now())

		timeline := &desc.Timeline{
			UserId: uint64(1),
			Type:   uint64(1),
			From:   &from,
			To:     &to,
		}

		Context("add", func() {
			It("correct date from/to", func() {
				mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Times(1)

				input := &desc.CreateTimelineV1Request{
					Timeline: timeline,
				}

				server.CreateTimelineV1(ctx, input)
			})
		})

		Context("get", func() {
			It("correct data", func() {
				result := &models.Timeline{
					Id:     1,
					UserId: 2,
					Type:   3,
					From:   utils.ConvertTimeInTimeStamp(time.Now()),
					To:     utils.ConvertTimeInTimeStamp(time.Now()),
				}

				mockRepo.EXPECT().DescribeEntity(ctx, gomock.Any()).Times(1).Return(result, nil)

				input := desc.GetTimelineV1Request{Id: 1}

				server.GetTimelineV1(ctx, &input)
			})

			It("not correct data = return error", func() {

				mockRepo.EXPECT().DescribeEntity(ctx, gomock.Any()).Times(1).Return(nil, errors.New("not found id"))

				input := desc.GetTimelineV1Request{Id: 1}

				_, err := server.GetTimelineV1(ctx, &input)

				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})

		Context("update", func() {
			It("correct data", func() {

				mockRepo.EXPECT().UpdateEntity(ctx, gomock.Any()).Times(1).Return(true, nil)

				input := &desc.UpdateTimelineV1Request{
					Timeline: timeline,
				}

				server.UpdateTimelineV1(ctx, input)
			})

			It("not correct data = return error", func() {

				mockRepo.EXPECT().UpdateEntity(ctx, gomock.Any()).Times(1).Return(false, errors.New("some want wro"))

				input := &desc.UpdateTimelineV1Request{
					Timeline: timeline,
				}

				_, err := server.UpdateTimelineV1(ctx, input)

				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})

		Context("list", func() {
			It("correct data", func() {

				mockRepo.EXPECT().ListEntities(ctx, gomock.Any(), gomock.Any()).Times(1)

				input := &desc.ListTimelineV1Request{
					Limit:  1,
					Offset: 0,
				}

				server.ListTimelineV1(ctx, input)
			})

			It("not correct data = return error", func() {

				mockRepo.EXPECT().ListEntities(ctx, gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, errors.New("some error"))

				input := &desc.ListTimelineV1Request{
					Limit:  1,
					Offset: 0,
				}

				_, err := server.ListTimelineV1(ctx, input)

				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})

		Context("remove", func() {
			It("correct data", func() {

				mockRepo.EXPECT().RemoveEntity(ctx, gomock.Any()).Times(1)

				input := &desc.RemoveTimelineV1Request{
					Id: 1,
				}

				server.RemoveTimelineV1(ctx, input)
			})

			It("not correct data = return error", func() {

				mockRepo.EXPECT().RemoveEntity(ctx, gomock.Any()).
					Times(1).Return(errors.New("some error"))

				input := &desc.RemoveTimelineV1Request{
					Id: 1,
				}

				_, err := server.RemoveTimelineV1(ctx, input)

				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
})
