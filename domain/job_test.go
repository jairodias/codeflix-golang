package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV1().String()
	video.FilePath = "valid_path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("valid_path", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
