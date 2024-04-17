package provisioning

import (
	"os"
	"testing"

	ranchFrame "github.com/rancher/shepherd/pkg/config"
	"github.com/rancher/tfp-automation/config"
	"github.com/rancher/tfp-automation/defaults/configs"
	set "github.com/rancher/tfp-automation/framework/set/provisioning"
	"github.com/sirupsen/logrus"
)

// BuildModule is a function that builds the Terraform module.
func BuildModule(t *testing.T) error {
	clusterConfig := new(config.TerratestConfig)
	ranchFrame.LoadConfig(configs.Terratest, clusterConfig)

	keyPath := set.SetKeyPath()

	err := set.SetConfigTF(clusterConfig, "", "")
	if err != nil {
		return err
	}

	module, err := os.ReadFile(keyPath + configs.MainTF)
	if err != nil {
		logrus.Errorf("Failed to read main.tf file contents. Error: %v", err)

		return err
	}

	t.Log(string(module))

	return nil
}
