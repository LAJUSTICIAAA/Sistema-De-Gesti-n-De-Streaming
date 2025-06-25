package services

import (
	"github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming/models"
)

var videos []models.Video

func GetVideos() []models.Video {
	return videos
}

func AddVideo(video models.Video) {
	videos = append(videos, video)
}
