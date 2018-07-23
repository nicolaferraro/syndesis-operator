package action

import (
	"errors"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	"github.com/syndesisio/syndesis-operator/pkg/apis/syndesis/v1alpha1"
	"github.com/syndesisio/syndesis-operator/pkg/syndesis/common"
	"github.com/syndesisio/syndesis-operator/pkg/syndesis/configuration"
)

// Upgrade a legacy Syndesis installation (installed with template) using the operator.
type UpgradeLegacy struct {}


func (a *UpgradeLegacy) CanExecute(syndesis *v1alpha1.Syndesis) bool {
	return syndesisInstallationStatusIs(syndesis, v1alpha1.SyndesisInstallationStatusUpgradingLegacy)
}

func (a *UpgradeLegacy) Execute(syndesis *v1alpha1.Syndesis) error {
	// Checking that there's only one installation to avoid stealing resources
	if anotherInstallation, err := isAnotherActiveInstallationPresent(syndesis); err != nil {
		return err
	} else if anotherInstallation {
		return errors.New("another syndesis installation active")
	}

	logrus.Info("Attaching Syndesis installation to resource ", syndesis.Name)

	err := common.AttachSyndesisToResource(syndesis)
	if err != nil {
		return err
	}

	syndesisVersion, err := configuration.GetSyndesisVersionFromNamespace(syndesis.Namespace)
	if err != nil {
		return err
	}

	target := syndesis.DeepCopy()
	target.Status.InstallationStatus = v1alpha1.SyndesisInstallationStatusStarting
	target.Status.Reason = v1alpha1.SyndesisStatusReasonMissing
	target.Status.Description = ""
	target.Status.Version = syndesisVersion

	logrus.Info("Syndesis installation attached to resource ", syndesis.Name)
	return sdk.Update(target)
}

func isAnotherActiveInstallationPresent(syndesis *v1alpha1.Syndesis) (bool, error) {
	lst := v1alpha1.NewSyndesisList()
	err := sdk.List(syndesis.Namespace, lst)
	if err != nil {
		return false, err
	}

	for _, that := range lst.Items {
		if that.Name != syndesis.Name &&
			that.Status.InstallationStatus != v1alpha1.SyndesisInstallationStatusNotInstalled &&
			that.Status.InstallationStatus != v1alpha1.SyndesisStatusReasonMissing {
				return true, nil
		}
	}

	return false, nil
}