package octopusdeploy

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectGet(t *testing.T) {
	client := getFakeOctopusClient(t, "/api/projects/Projects-663", http.StatusOK, getProjectResponseJSON)
	project, err := client.Project.Get("Projects-663")

	assert.Nil(t, err)
	assert.Equal(t, "Canary .NET Core 2.0", project.Name)
}

const getProjectResponseJSON = `
{
  "Id": "Projects-663",
  "VariableSetId": "variableset-Projects-663",
  "DeploymentProcessId": "deploymentprocess-Projects-663",
  "DiscreteChannelRelease": false,
  "IncludedLibraryVariableSetIds": [
    "LibraryVariableSets-124",
    "LibraryVariableSets-41",
    "LibraryVariableSets-21",
    "LibraryVariableSets-81"
  ],
  "DefaultToSkipIfAlreadyInstalled": false,
  "TenantedDeploymentMode": "Untenanted",
  "DefaultGuidedFailureMode": "EnvironmentDefault",
  "VersioningStrategy": {
    "DonorPackageStepId": null,
    "Template": "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}"
  },
  "ReleaseCreationStrategy": {
    "ReleaseCreationPackageStepId": "",
    "ChannelId": null
  },
  "Templates": [],
  "AutoDeployReleaseOverrides": [],
  "Name": "Canary .NET Core 2.0",
  "Slug": "canary-net-core-2-0",
  "Description": "Deployment pipeline for the canary .NET Core 2.0 skeleton project",
  "IsDisabled": false,
  "ProjectGroupId": "ProjectGroups-261",
  "LifecycleId": "Lifecycles-21",
  "AutoCreateRelease": false,
  "ProjectConnectivityPolicy": {
    "SkipMachineBehavior": "SkipUnavailableMachines",
    "TargetRoles": [],
    "AllowDeploymentsToNoTargets": false
  },
  "Links": {
    "Self": "/api/projects/Projects-663",
    "Releases": "/api/projects/Projects-663/releases{/version}{?skip,take,searchByVersion}",
    "Channels": "/api/projects/Projects-663/channels{?skip,take,partialName}",
    "Triggers": "/api/projects/Projects-663/triggers{?skip,take,partialName}",
    "OrderChannels": "/api/projects/Projects-663/channels/order",
    "Variables": "/api/variables/variableset-Projects-663",
    "Progression": "/api/progression/Projects-663{?aggregate}",
    "DeploymentProcess": "/api/deploymentprocesses/deploymentprocess-Projects-663",
    "Web": "/app#/projects/Projects-663",
    "Logo": "/api/projects/Projects-663/logo?cb=2018.3.6"
  }
}`
