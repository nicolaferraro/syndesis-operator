package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SyndesisList struct {
	metav1.TypeMeta  `json:",inline"`
	metav1.ListMeta  `json:"metadata"`
	Items []Syndesis `json:"items"`
}

func NewSyndesisList() *SyndesisList {
	return &SyndesisList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: groupName + "/" + version,
			Kind: "Syndesis",
		},
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Syndesis struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata"`
	Spec   SyndesisSpec   `json:"spec"`
	Status SyndesisStatus `json:"status,omitempty"`
}

type SyndesisSpec struct {
	RouteHostName    		string      `json:"routeHostname,omitempty"`
	DemoData         		*bool       `json:"demoData,omitempty"`
	DeployIntegrations		*bool		`json:"deployIntegrations,omitempty"`
	ImageStreamNamespace	string		`json:"imageStreamNamespace,omitempty"`
	IntegrationLimit 		*int        `json:"integrationLimit,omitempty"`
	Registry 				string		`json:"registry,omitempty"`
	Components       		Components  `json:"components,omitempty"`
	OpenShiftConsoleUrl     string      `json:"openShiftConsoleUrl,omitempty"`

}

type SyndesisPhase string

const (
	SyndesisPhaseMissing				SyndesisPhase = ""
	SyndesisPhaseInstalling            	SyndesisPhase = "Installing"
	SyndesisPhaseUpgradingLegacy        SyndesisPhase = "UpgradingLegacy"
	SyndesisPhaseStarting              	SyndesisPhase = "Starting"
	SyndesisPhaseStartupFailed         	SyndesisPhase = "StartupFailed"
	SyndesisPhaseInstalled            	SyndesisPhase = "Installed"
	SyndesisPhaseNotInstalled          	SyndesisPhase = "NotInstalled"
	SyndesisPhaseUpgrading             	SyndesisPhase = "Upgrading"
	SyndesisPhaseUpgradeFailureBackoff 	SyndesisPhase = "UpgradeFailureBackoff"
	SyndesisPhaseUpgradeFailed			SyndesisPhase = "UpgradeFailed"
)

type SyndesisStatusReason string

const (
	SyndesisStatusReasonMissing					= ""
	SyndesisStatusReasonDuplicate				= "Duplicate"
	SyndesisStatusReasonDeploymentNotReady		= "DeploymentNotReady"
	SyndesisStatusReasonUpgradePodFailed		= "UpgradePodFailed"
	SyndesisStatusReasonTooManyUpgradeAttempts	= "TooManyUpgradeAttempts"
)

type SyndesisStatus struct {
	Phase				SyndesisPhase				`json:"installationStatus,omitempty"`
	UpgradeAttempts		int32						`json:"upgradeAttempts,omitempty"`
	LastUpgradeFailure	*metav1.Time				`json:"lastUpgradeFailure,omitempty"`
	ForceUpgrade		bool						`json:"forceUpgrade,omitempty"`
	Reason				SyndesisStatusReason		`json:"reason,omitempty"`
	Description			string						`json:"description,omitempty"`
	Version				string						`json:"version,omitempty"`
	TargetVersion		string						`json:"targetVersion,omitempty"`
}


type Components struct {
	Db         DbConfiguration			`json:"db,omitempty"`
	Prometheus PrometheusConfiguration	`json:"prometheus,omitempty"`
	Server     ServerConfiguration		`json:"server,omitempty"`
	Meta       MetaConfiguration		`json:"meta,omitempty"`
}

type DbConfiguration struct {
	Resources 					ResourcesWithVolume		`json:"resources,omitempty"`
	User      					string                  `json:"user,omitempty"`
	Database    				string                  `json:"database,omitempty"`
	ImageStreamNamespace		string                  `json:"imageStreamNamespace,omitempty"`
}

type PrometheusConfiguration struct {
	Resources 					ResourcesWithVolume		`json:"resources,omitempty"`
}

type ServerConfiguration struct {
	Resources 					Resources				`json:"resources,omitempty"`
}

type MetaConfiguration struct {
	Resources 					ResourcesWithVolume		`json:"resources,omitempty"`
}

type Resources struct {
	v1.ResourceRequirements `json:",inline"`
}

type ResourcesWithVolume struct {
	v1.ResourceRequirements 				`json:",inline"`
	VolumeCapacity				string      `json:"volumeCapacity,omitempty"`
}
