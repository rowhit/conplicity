package providers

import (
	"testing"

	"github.com/docker/docker/api/types"
)

func TestMySQLGetName(t *testing.T) {
	expected := "MySQL"
	got := (&MySQLProvider{}).GetName()
	if expected != got {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestMySQLGetBackupDir(t *testing.T) {
	expected := "backups"
	got := (&MySQLProvider{}).GetBackupDir()
	if expected != got {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestMySQLGetPrepareCommand(t *testing.T) {
	mount := &types.MountPoint{
		Destination: "/mnt",
	}

	expected := []string{
		"sh",
		"-c",
		"mkdir -p /mnt/backups && mysqldump --all-databases --extended-insert --password=$MYSQL_ROOT_PASSWORD > /mnt/backups/all.sql",
	}
	got := (&MySQLProvider{}).GetPrepareCommand(mount)
	if len(got) != 3 {
		t.Fatalf("Expected command to have 3 elements, got %v", len(got))
	} else {
		if expected[2] != got[2] {
			t.Fatalf("Expected %s, got %s", expected, got)
		}
	}
}
