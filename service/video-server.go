package service

import "github.com/aashabtajwar/rest/entity"

type VideoService interface {
	// Save() method takes in a video and returns a video
	Save(entity.Video) entity.Video

	// returns a slice of videos
	FindAll() []entity.Video
}

// struct that stores a slice of videos
type videoService struct {
	videos []entity.Video
}

// constructor function
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
