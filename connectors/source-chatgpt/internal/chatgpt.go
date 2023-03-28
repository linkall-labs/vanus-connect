// Copyright 2023 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sashabaranov/go-openai"
)

const (
	responseEmpty = "Get ChatGPT response empty."
	responseErr   = "Get ChatGPT response failed."
)

var (
	ErrLimit = fmt.Errorf("reached the daily limit of using ChatGPT")
)

type chatGPTService struct {
	client       *openai.Client
	config       *chatGPTConfig
	lock         sync.Mutex
	messages     []openai.ChatCompletionMessage
	day          int
	num          int
	limitContent string
}

func newChatGPTService(config *chatGPTConfig) *chatGPTService {
	client := openai.NewClient(config.Token)
	return &chatGPTService{
		config:       config,
		client:       client,
		messages:     make([]openai.ChatCompletionMessage, 0, config.ContextCount),
		day:          today(),
		limitContent: fmt.Sprintf("You've reached the daily limit(%d/day) of using ChatGPT Source. Your quota will be restored tomorrow.", config.EverydayLimit),
	}
}

func today() int {
	return time.Now().UTC().Day()
}
func (s *chatGPTService) reset() {
	s.day = today()
	s.num = 0
	s.messages = make([]openai.ChatCompletionMessage, 0, s.config.ContextCount)
}

func (s *chatGPTService) CreateChatCompletion(content string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.num >= s.config.EverydayLimit {
		if today() == s.day {
			return s.limitContent, ErrLimit
		}
		s.reset()
	}
	if len(s.messages) >= s.config.ContextCount {
		index := len(s.messages) - s.config.ContextCount + 2
		s.messages = s.messages[index:]
	}
	s.messages = append(s.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: s.messages,
		},
	)
	if err != nil {
		return responseErr, err
	}
	s.num++
	if len(resp.Choices) == 0 {
		return responseEmpty, nil
	}
	respContent := resp.Choices[0].Message.Content
	if respContent == "" {
		return responseEmpty, nil
	}
	s.messages = append(s.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: respContent,
	})
	return respContent, nil
}
