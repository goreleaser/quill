package notary

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type StatusConfig struct {
	Timeout time.Duration
	Poll    time.Duration
	Wait    bool
}

func PollStatus(ctx context.Context, sub *Submission, cfg StatusConfig) (SubmissionStatus, error) {
	var err error

	ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	var status SubmissionStatus = PendingStatus

	for !status.isCompleted() {
		select {
		case <-ctx.Done():
			return "", errors.New("timeout waiting for notarize Submission response")

		default:
			status, err = sub.Status(ctx)
			if err != nil {
				return "", err
			}
		}

		time.Sleep(cfg.Poll)
	}

	if !status.isSuccessful() {
		logs, err := sub.Logs(ctx)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("submission result is %+v: %+v", status, logs)
	}

	return status, nil
}
