package main

import (
	"oadp/dm-plugin/plugin"

	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func main() {
	framework.NewServer().RegisterBackupItemAction("oadp/dm-plugin", newBackupPlugin)
}

func newBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewDMVolSyncPlugin(logger), nil
}
