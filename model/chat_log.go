package model

import (
	"context"
	"fmt"
	"github.com/songquanpeng/one-api/common/logger"
)

type ChatLog struct {
	Id       int    `json:"id" gorm:"primaryKey,autoIncrement"`
	UserId   int    `json:"user_id" gorm:"index"`
	Token    string `json:"created_at" gorm:"type:varchar(192);default:''"`
	Request  string `json:"request" gorm:"type:longtext"`
	Response string `json:"response" gorm:"type:longtext"`
	Model    string `json:"model" gorm:"type:varchar(64);default:''"`
}

func SaveChatLogRequest(ctx context.Context, UserID int, Request string, model string) int {
	value := &ChatLog{
		UserId:   UserID,
		Request:  Request,
		Response: "",
		Model:    model,
	}
	err := DB.Create(value).Error
	if err != nil {
		logger.Debug(ctx, fmt.Sprintf("save chat log error: %s %+v", err.Error(), value))
	}
	return value.Id
}

func SaveChatLogResponse(ctx context.Context, ChatLogId int, Response string) {
	if ChatLogId == 0 {
		return
	}
	value := &ChatLog{
		Id:       ChatLogId,
		Response: Response,
	}
	err := DB.Model(value).Updates(value).Error
	if err != nil {
		logger.Debug(ctx, fmt.Sprintf("save chat log error: %s %+v", err.Error(), value))
	}
}
