package backup

import (
	"time"

	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	// Maximum amount of time to wait for Backup changes to propagate
	propagationTimeout = 2 * time.Minute
	// Maximum amount of time to wait for Framework creation and deletion
	frameworkCreationTimeout = 2 * time.Minute
	frameworkDeletionTimeout = 2 * time.Minute
)

func waitFrameworkCreated(conn *backup.Backup, id string) (*backup.DescribeFrameworkOutput, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{frameworkStatusCreationInProgress},
		Target:  []string{frameworkStatusCompleted, frameworkStatusFailed},
		Refresh: statusFramework(conn, id),
		Timeout: frameworkCreationTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*backup.DescribeFrameworkOutput); ok {
		return output, err
	}

	return nil, err
}

func waitFrameworkDeleted(conn *backup.Backup, id string) (*backup.DescribeFrameworkOutput, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{frameworkStatusDeletionInProgress},
		Target:  []string{backup.ErrCodeResourceNotFoundException},
		Refresh: statusFramework(conn, id),
		Timeout: frameworkDeletionTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*backup.DescribeFrameworkOutput); ok {
		return output, err
	}

	return nil, err
}
