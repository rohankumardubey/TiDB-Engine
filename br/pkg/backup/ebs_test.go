// Copyright 2022 PingCAP, Inc. Licensed under Apache-2.0.

package backup

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	cfg := &EBSBackupInfo{}
	curDir, err := os.Getwd()
	require.NoError(t, err)
	require.NoError(t, cfg.ConfigFromFile(filepath.Join(curDir, "ebs_backup.json")))
}
