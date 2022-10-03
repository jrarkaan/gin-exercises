package service

import (
	"github.com/jrarkaan/gin-exercises/entity"
	"strings"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}
type responseSave struct {
	status  string
	message string
	videos  []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	for _, val := range service.videos {
		if strings.ToLower(strings.ReplaceAll(val.Title, " ", "")) == strings.ToLower(strings.ReplaceAll(video.Title, " ", "")) {

		}
	}
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
