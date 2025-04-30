package tests

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/armanbektassov/go_chat/internal/api/message"
	"github.com/armanbektassov/go_chat/internal/model"
	"github.com/armanbektassov/go_chat/internal/service"
	serviceMocks "github.com/armanbektassov/go_chat/internal/service/mocks"
	desc "github.com/armanbektassov/go_chat/pkg/message_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId  = gofakeit.Int64()
		creator = gofakeit.Animal()
		text    = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error")

		req = &desc.SendMessageRequest{
			MessageInfo: &desc.MessageInfo{
				ChatId:  chatId,
				Creator: creator,
				Text:    text,
			},
		}

		info = &model.MessageInfo{
			ChatId:  chatId,
			Creator: creator,
			Text:    text,
		}

		res = &emptypb.Empty{}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, info).Return(res, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, info).Return(res, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			messageServiceMock := tt.messageServiceMock(mc)
			api := message.NewImplementation(messageServiceMock)

			newID, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
