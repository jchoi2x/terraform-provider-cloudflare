// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectResultEnvelope struct {
	Result PagesProjectModel `json:"result"`
}

type PagesProjectModel struct {
	ID                  types.String                                                   `tfsdk:"id" json:"-,computed"`
	Name                types.String                                                   `tfsdk:"name" json:"name"`
	AccountID           types.String                                                   `tfsdk:"account_id" path:"account_id"`
	ProductionBranch    types.String                                                   `tfsdk:"production_branch" json:"production_branch"`
	BuildConfig         *PagesProjectBuildConfigModel                                  `tfsdk:"build_config" json:"build_config"`
	DeploymentConfigs   *PagesProjectDeploymentConfigsModel                            `tfsdk:"deployment_configs" json:"deployment_configs"`
	CreatedOn           timetypes.RFC3339                                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Subdomain           types.String                                                   `tfsdk:"subdomain" json:"subdomain,computed"`
	Domains             types.List                                                     `tfsdk:"domains" json:"domains,computed"`
	CanonicalDeployment customfield.NestedObject[PagesProjectCanonicalDeploymentModel] `tfsdk:"canonical_deployment" json:"canonical_deployment,computed"`
	LatestDeployment    customfield.NestedObject[PagesProjectLatestDeploymentModel]    `tfsdk:"latest_deployment" json:"latest_deployment,computed"`
	Source              customfield.NestedObject[PagesProjectSourceModel]              `tfsdk:"source" json:"source,computed"`
}

type PagesProjectBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed_optional"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed_optional"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed_optional"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed_optional"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed_optional"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed_optional"`
}

type PagesProjectDeploymentConfigsModel struct {
	Preview    customfield.NestedObject[PagesProjectDeploymentConfigsPreviewModel]    `tfsdk:"preview" json:"preview,computed_optional"`
	Production customfield.NestedObject[PagesProjectDeploymentConfigsProductionModel] `tfsdk:"production" json:"production,computed_optional"`
}

type PagesProjectDeploymentConfigsPreviewModel struct {
	AIBindings              map[string]jsontypes.Normalized                                              `tfsdk:"ai_bindings" json:"ai_bindings,computed_optional"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                                              `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed_optional"`
	Browsers                map[string]jsontypes.Normalized                                              `tfsdk:"browsers" json:"browsers,computed_optional"`
	CompatibilityDate       types.String                                                                 `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags      types.List                                                                   `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	D1Databases             map[string]jsontypes.Normalized                                              `tfsdk:"d1_databases" json:"d1_databases,computed_optional"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                                              `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed_optional"`
	EnvVars                 map[string]jsontypes.Normalized                                              `tfsdk:"env_vars" json:"env_vars,computed_optional"`
	HyperdriveBindings      map[string]jsontypes.Normalized                                              `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed_optional"`
	KVNamespaces            map[string]jsontypes.Normalized                                              `tfsdk:"kv_namespaces" json:"kv_namespaces,computed_optional"`
	MTLSCertificates        map[string]jsontypes.Normalized                                              `tfsdk:"mtls_certificates" json:"mtls_certificates,computed_optional"`
	Placement               customfield.NestedObject[PagesProjectDeploymentConfigsPreviewPlacementModel] `tfsdk:"placement" json:"placement,computed_optional"`
	QueueProducers          map[string]jsontypes.Normalized                                              `tfsdk:"queue_producers" json:"queue_producers,computed_optional"`
	R2Buckets               map[string]jsontypes.Normalized                                              `tfsdk:"r2_buckets" json:"r2_buckets,computed_optional"`
	Services                map[string]jsontypes.Normalized                                              `tfsdk:"services" json:"services,computed_optional"`
	VectorizeBindings       map[string]jsontypes.Normalized                                              `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed_optional"`
}

type PagesProjectDeploymentConfigsPreviewPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed_optional"`
}

type PagesProjectDeploymentConfigsProductionModel struct {
	AIBindings              map[string]jsontypes.Normalized                                                 `tfsdk:"ai_bindings" json:"ai_bindings,computed_optional"`
	AnalyticsEngineDatasets map[string]jsontypes.Normalized                                                 `tfsdk:"analytics_engine_datasets" json:"analytics_engine_datasets,computed_optional"`
	Browsers                map[string]jsontypes.Normalized                                                 `tfsdk:"browsers" json:"browsers,computed_optional"`
	CompatibilityDate       types.String                                                                    `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags      types.List                                                                      `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	D1Databases             map[string]jsontypes.Normalized                                                 `tfsdk:"d1_databases" json:"d1_databases,computed_optional"`
	DurableObjectNamespaces map[string]jsontypes.Normalized                                                 `tfsdk:"durable_object_namespaces" json:"durable_object_namespaces,computed_optional"`
	EnvVars                 map[string]jsontypes.Normalized                                                 `tfsdk:"env_vars" json:"env_vars,computed_optional"`
	HyperdriveBindings      map[string]jsontypes.Normalized                                                 `tfsdk:"hyperdrive_bindings" json:"hyperdrive_bindings,computed_optional"`
	KVNamespaces            map[string]jsontypes.Normalized                                                 `tfsdk:"kv_namespaces" json:"kv_namespaces,computed_optional"`
	MTLSCertificates        map[string]jsontypes.Normalized                                                 `tfsdk:"mtls_certificates" json:"mtls_certificates,computed_optional"`
	Placement               customfield.NestedObject[PagesProjectDeploymentConfigsProductionPlacementModel] `tfsdk:"placement" json:"placement,computed_optional"`
	QueueProducers          map[string]jsontypes.Normalized                                                 `tfsdk:"queue_producers" json:"queue_producers,computed_optional"`
	R2Buckets               map[string]jsontypes.Normalized                                                 `tfsdk:"r2_buckets" json:"r2_buckets,computed_optional"`
	Services                map[string]jsontypes.Normalized                                                 `tfsdk:"services" json:"services,computed_optional"`
	VectorizeBindings       map[string]jsontypes.Normalized                                                 `tfsdk:"vectorize_bindings" json:"vectorize_bindings,computed_optional"`
}

type PagesProjectDeploymentConfigsProductionPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed_optional"`
}

type PagesProjectCanonicalDeploymentModel struct {
	ID                types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Aliases           types.List                                                                      `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       customfield.NestedObject[PagesProjectCanonicalDeploymentBuildConfigModel]       `tfsdk:"build_config" json:"build_config,computed_optional"`
	CreatedOn         timetypes.RFC3339                                                               `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentTrigger customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                                 `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                    `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                      `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectCanonicalDeploymentLatestStageModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                               `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProjectID         types.String                                                                    `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                    `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                    `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectCanonicalDeploymentSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectCanonicalDeploymentStagesModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                    `tfsdk:"url" json:"url,computed"`
}

type PagesProjectCanonicalDeploymentBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed_optional"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed_optional"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed_optional"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed_optional"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed_optional"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed_optional"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerModel struct {
	Metadata customfield.NestedObject[PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel] `tfsdk:"metadata" json:"metadata,computed_optional"`
	Type     types.String                                                                            `tfsdk:"type" json:"type,computed"`
}

type PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectCanonicalDeploymentLatestStageModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectCanonicalDeploymentSourceModel struct {
	Config customfield.NestedObject[PagesProjectCanonicalDeploymentSourceConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	Type   types.String                                                               `tfsdk:"type" json:"type,computed_optional"`
}

type PagesProjectCanonicalDeploymentSourceConfigModel struct {
	DeploymentsEnabled           types.Bool   `tfsdk:"deployments_enabled" json:"deployments_enabled,computed_optional"`
	Owner                        types.String `tfsdk:"owner" json:"owner,computed_optional"`
	PathExcludes                 types.List   `tfsdk:"path_excludes" json:"path_excludes,computed_optional"`
	PathIncludes                 types.List   `tfsdk:"path_includes" json:"path_includes,computed_optional"`
	PrCommentsEnabled            types.Bool   `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed_optional"`
	PreviewBranchExcludes        types.List   `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed_optional"`
	PreviewBranchIncludes        types.List   `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed_optional"`
	PreviewDeploymentSetting     types.String `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed_optional"`
	ProductionBranch             types.String `tfsdk:"production_branch" json:"production_branch,computed_optional"`
	ProductionDeploymentsEnabled types.Bool   `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed_optional"`
	RepoName                     types.String `tfsdk:"repo_name" json:"repo_name,computed_optional"`
}

type PagesProjectCanonicalDeploymentStagesModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentModel struct {
	ID                types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Aliases           types.List                                                                   `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       customfield.NestedObject[PagesProjectLatestDeploymentBuildConfigModel]       `tfsdk:"build_config" json:"build_config,computed_optional"`
	CreatedOn         timetypes.RFC3339                                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	DeploymentTrigger customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                              `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                                 `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                                   `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectLatestDeploymentLatestStageModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProjectID         types.String                                                                 `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                                 `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                                 `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectLatestDeploymentSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectLatestDeploymentStagesModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                                 `tfsdk:"url" json:"url,computed"`
}

type PagesProjectLatestDeploymentBuildConfigModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching,computed_optional"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command,computed_optional"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir,computed_optional"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir,computed_optional"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag,computed_optional"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token,computed_optional"`
}

type PagesProjectLatestDeploymentDeploymentTriggerModel struct {
	Metadata customfield.NestedObject[PagesProjectLatestDeploymentDeploymentTriggerMetadataModel] `tfsdk:"metadata" json:"metadata,computed_optional"`
	Type     types.String                                                                         `tfsdk:"type" json:"type,computed"`
}

type PagesProjectLatestDeploymentDeploymentTriggerMetadataModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectLatestDeploymentLatestStageModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectLatestDeploymentSourceModel struct {
	Config customfield.NestedObject[PagesProjectLatestDeploymentSourceConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	Type   types.String                                                            `tfsdk:"type" json:"type,computed_optional"`
}

type PagesProjectLatestDeploymentSourceConfigModel struct {
	DeploymentsEnabled           types.Bool   `tfsdk:"deployments_enabled" json:"deployments_enabled,computed_optional"`
	Owner                        types.String `tfsdk:"owner" json:"owner,computed_optional"`
	PathExcludes                 types.List   `tfsdk:"path_excludes" json:"path_excludes,computed_optional"`
	PathIncludes                 types.List   `tfsdk:"path_includes" json:"path_includes,computed_optional"`
	PrCommentsEnabled            types.Bool   `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed_optional"`
	PreviewBranchExcludes        types.List   `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed_optional"`
	PreviewBranchIncludes        types.List   `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed_optional"`
	PreviewDeploymentSetting     types.String `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed_optional"`
	ProductionBranch             types.String `tfsdk:"production_branch" json:"production_branch,computed_optional"`
	ProductionDeploymentsEnabled types.Bool   `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed_optional"`
	RepoName                     types.String `tfsdk:"repo_name" json:"repo_name,computed_optional"`
}

type PagesProjectLatestDeploymentStagesModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectSourceModel struct {
	Config customfield.NestedObject[PagesProjectSourceConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	Type   types.String                                            `tfsdk:"type" json:"type,computed_optional"`
}

type PagesProjectSourceConfigModel struct {
	DeploymentsEnabled           types.Bool   `tfsdk:"deployments_enabled" json:"deployments_enabled,computed_optional"`
	Owner                        types.String `tfsdk:"owner" json:"owner,computed_optional"`
	PathExcludes                 types.List   `tfsdk:"path_excludes" json:"path_excludes,computed_optional"`
	PathIncludes                 types.List   `tfsdk:"path_includes" json:"path_includes,computed_optional"`
	PrCommentsEnabled            types.Bool   `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled,computed_optional"`
	PreviewBranchExcludes        types.List   `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes,computed_optional"`
	PreviewBranchIncludes        types.List   `tfsdk:"preview_branch_includes" json:"preview_branch_includes,computed_optional"`
	PreviewDeploymentSetting     types.String `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting,computed_optional"`
	ProductionBranch             types.String `tfsdk:"production_branch" json:"production_branch,computed_optional"`
	ProductionDeploymentsEnabled types.Bool   `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled,computed_optional"`
	RepoName                     types.String `tfsdk:"repo_name" json:"repo_name,computed_optional"`
}
