package service

import (
	"github.com/antonioroque200OK/chatservice/internal/infra/grpc/pb"
	"github.com/antonioroque200OK/chatservice/internal/usecase/chatcompletionstream"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
	ChatConfigStream            chatcompletionstream.ChatCompletionConfigInputDto
	StreamChannel               chan chatcompletionstream.ChatCompletionOutputDto
}

func NewChatService(chatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase, chatConfigStream chatcompletionstream.ChatCompletionConfigInputDto, streamChannel chan chatcompletionstream.ChatCompletionOutputDto) *ChatService {
	return &ChatService{
		ChatCompletionStreamUseCase: chatCompletionStreamUseCase,
		ChatConfigStream:            chatConfigStream,
		StreamChannel:               streamChannel,
	}
}

func (c *ChatService) ChatStream(req *pb.ChatRequest, stream pb.ChatService_ChatStreamServer) error {
	chatConfig := chatcompletionstream.ChatCompletionConfigInputDto{
		Model:               c.ChatConfigStream.Model,
		ModelMaxTokens:      c.ChatConfigStream.ModelMaxTokens,
		Temperature:         c.ChatConfigStream.Temperature,
		TopP:                c.ChatConfigStream.TopP,
		N:                   c.ChatConfigStream.N,
		Stop:                c.ChatConfigStream.Stop,
		MaxTokens:           c.ChatConfigStream.MaxTokens,
		InitalSystemMessage: c.ChatConfigStream.InitalSystemMessage,
	}
	input := chatcompletionstream.ChatCompletionInputDto{
		UserMessage: req.GetUserMessage(),
		UserID:      req.GetUserId(),
		ChatID:      req.GetChatId(),
		Config:      chatConfig,
	}

	ctx := stream.Context()
	go func() {
		for msg := range c.StreamChannel {
			stream.Send(&pb.ChatResponse{
				ChatId:  msg.ChatID,
				UserId:  msg.UserID,
				Content: msg.Content,
			})
		}
	}()

	_, err := c.ChatCompletionStreamUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
