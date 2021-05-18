package schedule

import (
	"context"
	"fmt"
	"time"

	scheduler "cloud.google.com/go/scheduler/apiv1beta1"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1beta1"
)

type Client struct {
	project  string
	location string
	service  *scheduler.CloudSchedulerClient
}

func New(ctx context.Context, project, location string) (*Client, error) {
	c, err := scheduler.NewCloudSchedulerClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{project: project, location: location, service: c}, nil
}

type Job struct {
	LastAttemptTime time.Time
}

func (c *Client) GetJob(ctx context.Context, jobID string) (*Job, error) {
	job, err := c.service.GetJob(ctx, &schedulerpb.GetJobRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/jobs/%s", c.project, c.location, jobID),
	})
	if err != nil {
		return nil, err
	}

	return &Job{LastAttemptTime: job.LastAttemptTime.AsTime()}, nil
}
