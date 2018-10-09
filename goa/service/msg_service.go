package service

import "goa/goa/repo"

type MsgService interface {
	// todo
}

type msgService struct {
	msgRepo     repo.MsgRepository
	msgBodyRepo repo.MsgBodyRepository
}

func NewMsgService(r repo.MsgRepository, r1 repo.MsgBodyRepository) MsgService {
	return &msgService{r, r1}
}
