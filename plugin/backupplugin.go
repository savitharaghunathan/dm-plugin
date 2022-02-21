package plugin

import (
	"github.com/sirupsen/logrus"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

type DMVolSyncPlugin struct {
	log logrus.FieldLogger
}

func NewDMVolSyncPlugin(log logrus.FieldLogger) *DMVolSyncPlugin {
	return &DMVolSyncPlugin{log: log}
}

func (d *DMVolSyncPlugin) AppliesTo() (velero.ResourceSelector, error) {
	d.log.Debug("DMVolSyncPlugin AppliesTo")

	return velero.ResourceSelector{
		IncludedResources: []string{"persistentvolumeclaims"},
	}, nil
}

func (d *DMVolSyncPlugin) Execute(item runtime.Unstructured, backup *v1.Backup) (runtime.Unstructured, []velero.ResourceIdentifier, error) {

	d.log.Info("Hello from my DMVolSyncPlugin!")

	metadata, err := meta.Accessor(item)
	if err != nil {
		return nil, nil, err
	}

	annotations := metadata.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}

	annotations["velero.io/dm-backup-plugin_type"] = "volsync"

	metadata.SetAnnotations(annotations)

	return item, nil, nil
}
