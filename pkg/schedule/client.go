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
	ExecTime time.Time
}

func (c *Client) PrevJob(ctx context.Context) (*Job, error) {
	it := c.service.ListJobs(ctx, &schedulerpb.ListJobsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", c.project, c.location),
	})
	job, err := it.Next()
	if err != nil {
		return nil, err
	}

	return &Job{ExecTime: job.LastAttemptTime.AsTime()}, nil
}
