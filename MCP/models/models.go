package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// MonitorPkgManager represents the MonitorPkgManager schema from the OpenAPI specification
type MonitorPkgManager struct {
	Name string `json:"name"` // Package manager name.
	Repositories []interface{} `json:"repositories,omitempty"` // A list of package repositories (i.e. maven-central, or npm) that defaults to the canonical package registry for the given package manager.
}

// Projectsnapshotsfilters represents the Projectsnapshotsfilters schema from the OpenAPI specification
type Projectsnapshotsfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// MonitorGraphDependency represents the MonitorGraphDependency schema from the OpenAPI specification
type MonitorGraphDependency struct {
	Nodeid string `json:"nodeId"` // Node id unique across the graph.
}

// Aggregatedprojectissues represents the Aggregatedprojectissues schema from the OpenAPI specification
type Aggregatedprojectissues struct {
	Issues []map[string]interface{} `json:"issues,omitempty"` // An array of identified issues
}

// IssuesFilters represents the IssuesFilters schema from the OpenAPI specification
type IssuesFilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// Ignore represents the Ignore schema from the OpenAPI specification
type Ignore struct {
	Ignorepath map[string]interface{} `json:"ignorePath"` // The path that should be ignored. Wildcards can be specified with a `*`.
}

// Projectsettings represents the Projectsettings schema from the OpenAPI specification
type Projectsettings struct {
	Autoremediationprs AutoRemediationPrs `json:"autoRemediationPrs,omitempty"`
	Pullrequestfailonlyforhighseverity bool `json:"pullRequestFailOnlyForHighSeverity,omitempty"` // If set to `true`, fail Snyk Test only for high and critical severity vulnerabilities
	Autodepupgradeignoreddependencies []interface{} `json:"autoDepUpgradeIgnoredDependencies,omitempty"` // An array of comma-separated strings with names of dependencies you wish Snyk to ignore to upgrade.
	Autodepupgradelimit float64 `json:"autoDepUpgradeLimit,omitempty"` // The limit on auto dependency upgrade PRs.
	Autodepupgrademinage float64 `json:"autoDepUpgradeMinAge,omitempty"` // The age (in days) that an automatic dependency check is valid for
	Pullrequestassignment PullRequestAssignment `json:"pullRequestAssignment,omitempty"`
	Pullrequestfailonanyvulns bool `json:"pullRequestFailOnAnyVulns,omitempty"` // If set to `true`, fail Snyk Test if the repo has any vulnerabilities. Otherwise, fail only when the PR is adding a vulnerable dependency.
	Pullrequesttestenabled bool `json:"pullRequestTestEnabled,omitempty"` // If set to `true`, Snyk Test checks PRs for vulnerabilities.:cq
	Autodepupgradeenabled bool `json:"autoDepUpgradeEnabled,omitempty"` // If set to `true`, Snyk will raise dependency upgrade PRs automatically.
}

// Integrationsbody represents the Integrationsbody schema from the OpenAPI specification
type Integrationsbody struct {
}

// Orgsettingsresponse represents the Orgsettingsresponse schema from the OpenAPI specification
type Orgsettingsresponse struct {
	Requestaccess map[string]interface{} `json:"requestAccess,omitempty"` // Will only be returned if `API_KEY` has read access to request access settings.
}

// Notificationsettingresponse represents the Notificationsettingresponse schema from the OpenAPI specification
type Notificationsettingresponse struct {
	Enabled bool `json:"enabled"` // Whether notifications should be sent
	Issueseverity string `json:"issueSeverity"` // The severity levels of issues to send notifications for (only applicable for `new-remediations-vulnerabilities` notificationType)
	Issuetype string `json:"issueType"` // Filter the types of issue to include in notifications (only applicable for `new-remediations-vulnerabilities` notificationType)
	Inherited bool `json:"inherited,omitempty"` // Whether the setting was found on the requested context directly or inherited from a parent
}

// MavenFile represents the MavenFile schema from the OpenAPI specification
type MavenFile struct {
	Contents string `json:"contents"` // The contents of the file, encoded according to the `encoding` field.
}

// GroupsAuditlogsfilters represents the GroupsAuditlogsfilters schema from the OpenAPI specification
type GroupsAuditlogsfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// Govendorrequestpayload represents the Govendorrequestpayload schema from the OpenAPI specification
type Govendorrequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// Integrations represents the Integrations schema from the OpenAPI specification
type Integrations struct {
	Key string `json:"key,omitempty"` // The name of an integration
	Value string `json:"value,omitempty"` // Alphanumeric UUID including - with a limit of 36 characters
}

// MonitorGraph represents the MonitorGraph schema from the OpenAPI specification
type MonitorGraph struct {
	Rootnodeid string `json:"rootNodeId"` // Root node id. Note the root node name is used as your project name.
	Nodes []interface{} `json:"nodes"` // Array of node objects.
}

// Projectsnapshots represents the Projectsnapshots schema from the OpenAPI specification
type Projectsnapshots struct {
	Total float64 `json:"total,omitempty"` // The total number of results
	Snapshots []map[string]interface{} `json:"snapshots,omitempty"` // A list of the project's snapshots, ordered according to date (latest first).
}

// Projectsfilters represents the Projectsfilters schema from the OpenAPI specification
type Projectsfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// MonitorMetaData represents the MonitorMetaData schema from the OpenAPI specification
type MonitorMetaData struct {
	Targetframework string `json:"targetFramework,omitempty"` // Required for a NuGet or Paket DepGraph only. Specify the target framework in your project file using Target Framework Monikers (TFMs). For example, netstandard1.0, netcoreapp1.0 or net452. Test each framework separately if you have multiple defined.
}

// Graph represents the Graph schema from the OpenAPI specification
type Graph struct {
	Nodes []interface{} `json:"nodes"` // Array of node objects.
	Rootnodeid string `json:"rootNodeId"` // Root node id.
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Alias string `json:"alias,omitempty"` // deb, apk and rpm package managers should use an alias to indicate the target Operating System, for example 'debian:10'.
}

// Groupsettings represents the Groupsettings schema from the OpenAPI specification
type Groupsettings struct {
	Requestaccess map[string]interface{} `json:"requestAccess,omitempty"` // Can only be updated if `API_KEY` has edit access to request access settings.
	Sessionlength float64 `json:"sessionLength,omitempty"` // The new session length for the group in minutes. This must be an integer between 1 and 43200 (30 days). Setting this value to null will result in this group inheriting from the global default of 30 days.
}

// MonitorNode represents the MonitorNode schema from the OpenAPI specification
type MonitorNode struct {
	Deps []interface{} `json:"deps,omitempty"` // An array of package ids this package depends on.
	Nodeid string `json:"nodeId"` // Node id unique across the graph.
	Pkgid string `json:"pkgId"` // Package id reference should match id in pkg array and take the format name@version.
}

// Projectmove represents the Projectmove schema from the OpenAPI specification
type Projectmove struct {
	Targetorgid string `json:"targetOrgId,omitempty"` // The ID of the organization that the project should be moved to. The API_KEY must have group admin permissions. If the project is moved to a new group, a personal level API key is needed.
}

// Addmemberbody represents the Addmemberbody schema from the OpenAPI specification
type Addmemberbody struct {
	Role string `json:"role,omitempty"` // The role of the user, "admin" or "collaborator".
	Userid string `json:"userId,omitempty"` // The id of the user.
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	From map[string]interface{} `json:"from,omitempty"` // Paths from which the vulnerable package is required in the code base.
	Modificationtime string `json:"modificationTime,omitempty"`
	Upgradepath map[string]interface{} `json:"upgradePath,omitempty"`
	Creationtime string `json:"creationTime,omitempty"`
	Cvssscore float64 `json:"cvssScore,omitempty"` // CVSS Score.
	Ispatchable bool `json:"isPatchable,omitempty"` // Is a patch by Snyk available to fix this vulnerability?
	Packagename string `json:"packageName,omitempty"` // The name of the vulnerable package.
	Publicationtime string `json:"publicationTime,omitempty"`
	Identifiers map[string]interface{} `json:"identifiers,omitempty"` // Additional identifiers for this issue (CVE, CWE, etc).
	Language string `json:"language,omitempty"` // The programming language for this package.
	Severity string `json:"severity,omitempty"` // Snyk severity for this issue. One of: `critical`, `medium`, `high`, `medium` or `low`.
	Packagemanager__npm string `json:"packageManager `npm`,omitempty"`
	Cvssv3 string `json:"CVSSv3,omitempty"` // Common Vulnerability Scoring System (CVSS) provides a way to capture the principal characteristics of a vulnerability, and produce a numerical score reflecting its severity, as well as a textual representation of that score.
	Functions []interface{} `json:"functions,omitempty"` // List of vulnerable functions inside the vulnerable packages.
	Patches []interface{} `json:"patches,omitempty"` // Patches to fix this issue, by Snyk.
	Semver SemverObject `json:"semver,omitempty"`
	Description string `json:"description,omitempty"` // The description of the vulnerability
	Ispinnable bool `json:"isPinnable,omitempty"` // Will pinning this package to a newer version fix the vulnerability?
	Title string `json:"title,omitempty"` // The title of the vulnerability
	Credit map[string]interface{} `json:"credit,omitempty"` // The reporter of the vulnerability
	Isupgradable bool `json:"isUpgradable,omitempty"` // Will upgrading a top-level dependency fix the vulnerability?
	Alternativeids map[string]interface{} `json:"alternativeIds,omitempty"`
	Exploitmaturity string `json:"exploitMaturity,omitempty"` // Snyk exploit maturity for this issue. One of: `mature`, `proof-of-concept`, `no-known-exploit` or `no-data`.
	Disclosuretime string `json:"disclosureTime,omitempty"`
}

// TestCounts represents the TestCounts schema from the OpenAPI specification
type TestCounts struct {
	Results []map[string]interface{} `json:"results"` // A list of test counts
}

// ProjectCounts represents the ProjectCounts schema from the OpenAPI specification
type ProjectCounts struct {
	Results []map[string]interface{} `json:"results"` // A list of project counts by day
}

// Licenses represents the Licenses schema from the OpenAPI specification
type Licenses struct {
	Results []map[string]interface{} `json:"results"` // A list of licenses
	Total float64 `json:"total,omitempty"` // The number of results returned
}

// Mavenrequestpayload represents the Mavenrequestpayload schema from the OpenAPI specification
type Mavenrequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Id string `json:"id"` // Unique package identifier, should take the format name@version.
	Info PackageInfo `json:"info"`
}

// Gradlerequestpayload represents the Gradlerequestpayload schema from the OpenAPI specification
type Gradlerequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// PkgManager represents the PkgManager schema from the OpenAPI specification
type PkgManager struct {
	Name string `json:"name"` // Package manager name.
	Repositories []interface{} `json:"repositories,omitempty"` // A list of package repositories (i.e. maven-central, or npm) that defaults to the canonical package registry for the given package manager.
}

// SBTFile represents the SBTFile schema from the OpenAPI specification
type SBTFile struct {
	Contents string `json:"contents"` // The contents of the file, encoded according to the `encoding` field.
}

// Issues represents the Issues schema from the OpenAPI specification
type Issues struct {
	Results []map[string]interface{} `json:"results"` // A list of issues
	Total float64 `json:"total"` // The total number of results found
}

// Projectattributes represents the Projectattributes schema from the OpenAPI specification
type Projectattributes struct {
	Criticality []interface{} `json:"criticality,omitempty"`
	Environment []interface{} `json:"environment,omitempty"`
	Lifecycle []interface{} `json:"lifecycle,omitempty"`
}

// Sbtrequestpayload represents the Sbtrequestpayload schema from the OpenAPI specification
type Sbtrequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Code float64 `json:"code"` // The error response code
	ErrorField map[string]interface{} `json:"error"`
	Ok bool `json:"ok"`
}

// Function represents the Function schema from the OpenAPI specification
type Function struct {
	Functionid FunctionId `json:"functionId,omitempty"`
	Version []interface{} `json:"version,omitempty"` // Versions this function relates to.
}

// MonitorDepGraphData represents the MonitorDepGraphData schema from the OpenAPI specification
type MonitorDepGraphData struct {
	Schemaversion string `json:"schemaVersion"` // Snyk DepGraph library schema version.
	Graph MonitorGraph `json:"graph"`
	Pkgmanager MonitorPkgManager `json:"pkgManager"`
	Pkgs []interface{} `json:"pkgs"` // Array of package dependencies.
}

// Yarnrequestpayload represents the Yarnrequestpayload schema from the OpenAPI specification
type Yarnrequestpayload struct {
	Files map[string]interface{} `json:"files"` // The manifest files:
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
}

// PriorityScore represents the PriorityScore schema from the OpenAPI specification
type PriorityScore struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}

// Projectissuesfilters represents the Projectissuesfilters schema from the OpenAPI specification
type Projectissuesfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// PackageInfo represents the PackageInfo schema from the OpenAPI specification
type PackageInfo struct {
	Name string `json:"name"` // Package name.
	Version string `json:"version"` // Package version.
}

// Allignores represents the Allignores schema from the OpenAPI specification
type Allignores struct {
	Issueid []interface{} `json:"issueId"` // The issue ID that should be ignored.
}

// MonitorPackageInfo represents the MonitorPackageInfo schema from the OpenAPI specification
type MonitorPackageInfo struct {
	Name string `json:"name"` // Package name.
	Version string `json:"version"` // Package version.
}

// Patch represents the Patch schema from the OpenAPI specification
type Patch struct {
	Id string `json:"id,omitempty"`
	Modificationtime string `json:"modificationTime,omitempty"`
	Urls []interface{} `json:"urls,omitempty"` // Links to patch files to fix an issue.
	Version string `json:"version,omitempty"` // Versions this patch is applicable to, in semver format.
	Comments []interface{} `json:"comments,omitempty"`
}

// MonitorPackage represents the MonitorPackage schema from the OpenAPI specification
type MonitorPackage struct {
	Id string `json:"id"` // Unique package identifier, should take the format name@version.
	Info MonitorPackageInfo `json:"info"`
}

// GoPkgLock represents the GoPkgLock schema from the OpenAPI specification
type GoPkgLock struct {
	Contents string `json:"contents,omitempty"`
}

// GraphDependency represents the GraphDependency schema from the OpenAPI specification
type GraphDependency struct {
	Nodeid string `json:"nodeId"` // Node id unique across the graph.
}

// Node represents the Node schema from the OpenAPI specification
type Node struct {
	Deps []interface{} `json:"deps,omitempty"` // An array of package ids this package depends on.
	Nodeid string `json:"nodeId"` // Node id unique across the graph.
	Pkgid string `json:"pkgId"` // Package id reference should match id in pkg array and take the format name@version.
}

// Notificationsettingsrequest represents the Notificationsettingsrequest schema from the OpenAPI specification
type Notificationsettingsrequest struct {
	Weekly_report Simplenotificationsettingrequest `json:"weekly-report,omitempty"`
	New_issues_remediations Newissuesnotificationsettingrequest `json:"new-issues-remediations,omitempty"`
	Project_imported Simplenotificationsettingrequest `json:"project-imported,omitempty"`
	Test_limit Simplenotificationsettingrequest `json:"test-limit,omitempty"`
}

// Projectdependencygraph represents the Projectdependencygraph schema from the OpenAPI specification
type Projectdependencygraph struct {
	Depgraph map[string]interface{} `json:"depGraph"` // The dependency-graph object
}

// Listallprojects represents the Listallprojects schema from the OpenAPI specification
type Listallprojects struct {
	Org map[string]interface{} `json:"org,omitempty"`
	Projects []Projectwithoutremediation `json:"projects,omitempty"` // A list of org's projects
}

// SemverObject represents the SemverObject schema from the OpenAPI specification
type SemverObject struct {
	Vulnerable string `json:"vulnerable,omitempty"` // The (semver) range of versions vulnerable to this issue.
	Unaffected string `json:"unaffected,omitempty"` // The (semver) range of versions NOT vulnerable to this issue. *Deprecated* and should not be used.
}

// IssueCountsFilters represents the IssueCountsFilters schema from the OpenAPI specification
type IssueCountsFilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// Golangdeprequestpayload represents the Golangdeprequestpayload schema from the OpenAPI specification
type Golangdeprequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// Composerrequestpayload represents the Composerrequestpayload schema from the OpenAPI specification
type Composerrequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// Dependencies represents the Dependencies schema from the OpenAPI specification
type Dependencies struct {
	Results []map[string]interface{} `json:"results"` // A list of issues
	Total float64 `json:"total,omitempty"` // The number of results returned
}

// Piprequestpayload represents the Piprequestpayload schema from the OpenAPI specification
type Piprequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// Ignorerule represents the Ignorerule schema from the OpenAPI specification
type Ignorerule struct {
	Reasontype string `json:"reasonType"` // The classification of the ignore.
	Disregardiffixable bool `json:"disregardIfFixable"` // Only ignore the issue if no upgrade or patch is available.
	Expires string `json:"expires,omitempty"` // The timestamp that the issue will no longer be ignored.
	Ignorepath string `json:"ignorePath,omitempty"` // The path to ignore (default is `*` which represents all paths).
	Reason string `json:"reason,omitempty"` // The reason that the issue was ignored.
}

// Jiraissue represents the Jiraissue schema from the OpenAPI specification
type Jiraissue struct {
	Jiraissue map[string]interface{} `json:"jiraIssue,omitempty"` // The details about the jira issue.
}

// Newissuesnotificationsettingrequest represents the Newissuesnotificationsettingrequest schema from the OpenAPI specification
type Newissuesnotificationsettingrequest struct {
	Enabled bool `json:"enabled"` // Whether notifications should be sent
	Issueseverity string `json:"issueSeverity"` // The severity levels of issues to send notifications for (only applicable for `new-remediations-vulnerabilities` notificationType)
	Issuetype string `json:"issueType"` // Filter the types of issue to include in notifications (only applicable for `new-remediations-vulnerabilities` notificationType)
}

// Dependenciesfilters represents the Dependenciesfilters schema from the OpenAPI specification
type Dependenciesfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// Notificationsettingsresponse represents the Notificationsettingsresponse schema from the OpenAPI specification
type Notificationsettingsresponse struct {
	Project_imported Simplenotificationsettingresponse `json:"project-imported,omitempty"`
	Test_limit Simplenotificationsettingresponse `json:"test-limit,omitempty"`
	Weekly_report Simplenotificationsettingresponse `json:"weekly-report,omitempty"`
	New_issues_remediations Notificationsettingresponse `json:"new-issues-remediations,omitempty"`
}

// Graphrequestpayload represents the Graphrequestpayload schema from the OpenAPI specification
type Graphrequestpayload struct {
	Depgraph DepGraphData `json:"depGraph"`
}

// IntegrationCredentials represents the IntegrationCredentials schema from the OpenAPI specification
type IntegrationCredentials struct {
}

// Aggregatedprojectissuesfilters represents the Aggregatedprojectissuesfilters schema from the OpenAPI specification
type Aggregatedprojectissuesfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
	Includedescription bool `json:"includeDescription,omitempty"` // If set to `true`, Include issue's description, if set to `false` (by default), it won't (Non-IaC projects only)
	Includeintroducedthrough bool `json:"includeIntroducedThrough,omitempty"` // If set to `true`, Include issue's introducedThrough, if set to `false` (by default), it won't. It's for container only projects (Non-IaC projects only)
}

// MavenAdditionalFile represents the MavenAdditionalFile schema from the OpenAPI specification
type MavenAdditionalFile struct {
	Contents string `json:"contents"` // The contents of the file, encoded according to the `encoding` field.
}

// Alljiraissues represents the Alljiraissues schema from the OpenAPI specification
type Alljiraissues struct {
	Issueid []interface{} `json:"issueId"` // The issue ID and relating jira issue.
}

// Monitorgraphpayload represents the Monitorgraphpayload schema from the OpenAPI specification
type Monitorgraphpayload struct {
	Depgraph MonitorDepGraphData `json:"depGraph"`
	Meta MonitorMetaData `json:"meta,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Targetreference string `json:"targetReference,omitempty"` // The identifier for which revision of the resource is scanned by Snyk. For example this may be a branch for SCM project, or a tag for a container image
	Readonly bool `json:"readOnly,omitempty"` // Whether the project is read-only
	TypeField string `json:"type,omitempty"` // The package manager of the project
	Attributes Projectattributes `json:"attributes,omitempty"`
	Imageplatform string `json:"imagePlatform,omitempty"` // For docker projects shows the platform of the image
	Importinguser map[string]interface{} `json:"importingUser,omitempty"` // The user who imported the project
	Id string `json:"id,omitempty"` // The project identifier
	Ismonitored bool `json:"isMonitored,omitempty"` // Describes if a project is currently monitored or it is de-activated
	Name string `json:"name,omitempty"`
	Lasttesteddate string `json:"lastTestedDate,omitempty"` // The date on which the most recent test was conducted for this project
	Branch string `json:"branch,omitempty"` // The monitored branch (if available)
	Browseurl string `json:"browseUrl,omitempty"` // URL with project overview
	Tags []interface{} `json:"tags,omitempty"` // List of applied tags
	Totaldependencies float64 `json:"totalDependencies,omitempty"` // Number of dependencies of the project
	Imagebaseimage string `json:"imageBaseImage,omitempty"` // For docker projects shows the base image
	Imagetag string `json:"imageTag,omitempty"` // For docker projects shows the tag of the image
	Remediation map[string]interface{} `json:"remediation,omitempty"` // Remediation data (if available)
	Remoterepourl string `json:"remoteRepoUrl,omitempty"` // The project remote repository url. Only set for projects imported via the Snyk CLI tool.
	Issuecountsbyseverity map[string]interface{} `json:"issueCountsBySeverity,omitempty"` // Number of known vulnerabilities in the project, not including ignored issues
	Created string `json:"created,omitempty"` // The date that the project was created on
	Owner map[string]interface{} `json:"owner,omitempty"` // The user who owns the project, null if not set { "id": "e713cf94-bb02-4ea0-89d9-613cce0caed2", "name": "example-user@snyk.io", "username": "exampleUser", "email": "example-user@snyk.io" }
	Hostname string `json:"hostname,omitempty"` // The hostname for a CLI project, null if not set
	Origin string `json:"origin,omitempty"` // The origin the project was added from
	Testfrequency string `json:"testFrequency,omitempty"` // The frequency of automated Snyk re-test. Can be 'daily', 'weekly or 'never'
	Imagecluster string `json:"imageCluster,omitempty"` // For Kubernetes projects shows the origin cluster name
	Imageid string `json:"imageId,omitempty"` // For docker projects shows the ID of the image
}

// Simplenotificationsettingresponse represents the Simplenotificationsettingresponse schema from the OpenAPI specification
type Simplenotificationsettingresponse struct {
	Enabled bool `json:"enabled"` // Whether notifications should be sent
	Inherited bool `json:"inherited,omitempty"` // Whether the setting was found on the requested context directly or inherited from a parent
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key string `json:"key,omitempty"` // Alphanumeric including - and _ with a limit of 30 characters
	Value string `json:"value,omitempty"` // Alphanumeric including - and _ with a limit of 50 characters
}

// Projectwithoutremediation represents the Projectwithoutremediation schema from the OpenAPI specification
type Projectwithoutremediation struct {
	Origin string `json:"origin,omitempty"` // The origin the project was added from
	Lasttesteddate string `json:"lastTestedDate,omitempty"` // The date on which the most recent test was conducted for this project
	TypeField string `json:"type,omitempty"` // The package manager of the project
	Testfrequency string `json:"testFrequency,omitempty"` // The frequency of automated Snyk re-test. Can be 'daily', 'weekly or 'never'
	Issuecountsbyseverity map[string]interface{} `json:"issueCountsBySeverity,omitempty"` // Number of known vulnerabilities in the project, not including ignored issues
	Remoterepourl string `json:"remoteRepoUrl,omitempty"` // The project remote repository url. Only set for projects imported via the Snyk CLI tool.
	Targetreference string `json:"targetReference,omitempty"` // The identifier for which revision of the resource is scanned by Snyk. For example this may be a branch for SCM project, or a tag for a container image
	Readonly bool `json:"readOnly,omitempty"` // Whether the project is read-only
	Browseurl string `json:"browseUrl,omitempty"` // URL with project overview
	Id string `json:"id,omitempty"` // The project identifier
	Imageplatform string `json:"imagePlatform,omitempty"` // For docker projects shows the platform of the image
	Imagetag string `json:"imageTag,omitempty"` // For docker projects shows the tag of the image
	Name string `json:"name,omitempty"`
	Branch string `json:"branch,omitempty"` // The monitored branch (if available)
	Imagebaseimage string `json:"imageBaseImage,omitempty"` // For docker projects shows the base image
	Owner map[string]interface{} `json:"owner,omitempty"` // The user who owns the project, null if not set { "id": "e713cf94-bb02-4ea0-89d9-613cce0caed2", "name": "example-user@snyk.io", "username": "exampleUser", "email": "example-user@snyk.io" }
	Importinguser map[string]interface{} `json:"importingUser,omitempty"` // The user who imported the project
	Tags []interface{} `json:"tags,omitempty"` // List of applied tags
	Totaldependencies float64 `json:"totalDependencies,omitempty"` // Number of dependencies of the project
	Created string `json:"created,omitempty"` // The date that the project was created on
	Imagecluster string `json:"imageCluster,omitempty"` // For Kubernetes projects shows the origin cluster name
	Ismonitored bool `json:"isMonitored,omitempty"` // Describes if a project is currently monitored or it is de-activated
	Attributes Projectattributes `json:"attributes,omitempty"`
	Imageid string `json:"imageId,omitempty"` // For docker projects shows the ID of the image
}

// PullRequestAssignment represents the PullRequestAssignment schema from the OpenAPI specification
type PullRequestAssignment struct {
	Assignees []interface{} `json:"assignees,omitempty"` // an array of usernames that have contributed to the organization's project(s).
	Enabled bool `json:"enabled,omitempty"` // if the organization's project(s) will assign Snyk pull requests.
	TypeField string `json:"type,omitempty"`
}

// TestsFilters represents the TestsFilters schema from the OpenAPI specification
type TestsFilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// AutoRemediationPrs represents the AutoRemediationPrs schema from the OpenAPI specification
type AutoRemediationPrs struct {
	Backlogprsenabled bool `json:"backlogPrsEnabled,omitempty"` // If true, allows automatic remediation of newly identified issues, or older issues where a fix has been identified
	Freshprsenabled bool `json:"freshPrsEnabled,omitempty"` // If true, allows automatic remediation of prioritized backlog issues
	Usepatchremediation bool `json:"usePatchRemediation,omitempty"` // If true, allows using patched remediation
}

// OrgAuditlogsfilters represents the OrgAuditlogsfilters schema from the OpenAPI specification
type OrgAuditlogsfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// BrokerSettings represents the BrokerSettings schema from the OpenAPI specification
type BrokerSettings struct {
	Enabled bool `json:"enabled,omitempty"`
}

// Createorganizationsbody represents the Createorganizationsbody schema from the OpenAPI specification
type Createorganizationsbody struct {
	Sourceorgid string `json:"sourceOrgId,omitempty"` // The id of an organization to copy settings from. If provided, this organization must be associated with the same group. The items that will be copied are: Source control integrations (GitHub, GitLab, BitBucket) \+ Container registries integrations (ACR, Docker Hub, ECR, GCR) \+ Container orchestrators integrations (Kubernetes) \+ PaaS and Serverless Integrations (Heroku, AWS Lambda) \+ Notification integrations (Slack, Jira) \+ Policies \+ Ignore settings \+ Language settings \+ Infrastructure as Code settings \+ Snyk Code settings The following will not be copied across: Service accounts \+ Members \+ Projects \+ Notification preferences
	Groupid string `json:"groupId,omitempty"` // The group ID. The `API_KEY` must have access to this group.
	Name string `json:"name"` // The name of the new organization
}

// YarnLockFile represents the YarnLockFile schema from the OpenAPI specification
type YarnLockFile struct {
	Contents string `json:"contents,omitempty"`
}

// Issuepaths represents the Issuepaths schema from the OpenAPI specification
type Issuepaths struct {
	Paths [][]map[string]interface{} `json:"paths,omitempty"` // A list of the dependency paths that introduce the issue
	Snapshotid string `json:"snapshotId,omitempty"` // The identifier of the snapshot for which the paths have been found
	Total float64 `json:"total,omitempty"` // The total number of results
	Links map[string]interface{} `json:"links,omitempty"` // Onward links from this record
}

// Integrationsettings represents the Integrationsettings schema from the OpenAPI specification
type Integrationsettings struct {
	Autodepupgradeignoreddependencies []interface{} `json:"autoDepUpgradeIgnoredDependencies,omitempty"` // A list of strings defining what dependencies should be ignored
	Autodepupgradelimit float64 `json:"autoDepUpgradeLimit,omitempty"` // A limit on how many automatic dependency upgrade PRs can be opened simultaneously
	Autodepupgradeenabled bool `json:"autoDepUpgradeEnabled,omitempty"` // Defines if the functionality is enabled
	Autodepupgrademinage float64 `json:"autoDepUpgradeMinAge,omitempty"` // The age (in days) that an automatic dependency check is valid for
	Manualremediationprs map[string]interface{} `json:"manualRemediationPrs,omitempty"` // Defines manual remediation policies
	Pullrequestassignment PullRequestAssignment `json:"pullRequestAssignment,omitempty"`
	Pullrequestfailonlyforhighseverity bool `json:"pullRequestFailOnlyForHighSeverity,omitempty"` // If an opened PR only should fail its validation if any dependencies are marked as being of high severity
	Pullrequesttestenabled bool `json:"pullRequestTestEnabled,omitempty"` // If opened PRs should be tested
	Autoremediationprs map[string]interface{} `json:"autoRemediationPrs,omitempty"` // Defines automatic remediation policies
	Dockerfilescmenabled bool `json:"dockerfileSCMEnabled,omitempty"` // If true, will automatically detect and scan Dockerfiles in your Git repositories, surface base image vulnerabilities and recommend possible fixes
	Pullrequestfailonanyvulns bool `json:"pullRequestFailOnAnyVulns,omitempty"` // If an opened PR should fail to be validated if any vulnerable dependencies have been detected
}

// Jiraissuerequest represents the Jiraissuerequest schema from the OpenAPI specification
type Jiraissuerequest struct {
	Fields map[string]interface{} `json:"fields,omitempty"`
}

// Orgsettingsrequest represents the Orgsettingsrequest schema from the OpenAPI specification
type Orgsettingsrequest struct {
	Requestaccess map[string]interface{} `json:"requestAccess,omitempty"` // Can only be updated if `API_KEY` has edit access to request access settings.
}

// Tagbody represents the Tagbody schema from the OpenAPI specification
type Tagbody struct {
	Key string `json:"key,omitempty"` // Valid tag key.
	Value string `json:"value,omitempty"` // Valid tag value.
}

// Simplenotificationsettingrequest represents the Simplenotificationsettingrequest schema from the OpenAPI specification
type Simplenotificationsettingrequest struct {
	Enabled bool `json:"enabled"` // Whether notifications should be sent
}

// Npmrequestpayload represents the Npmrequestpayload schema from the OpenAPI specification
type Npmrequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// PackageLockJsonFile represents the PackageLockJsonFile schema from the OpenAPI specification
type PackageLockJsonFile struct {
	Contents string `json:"contents,omitempty"` // The contents of the file, encoded according to the `encoding` field.
}

// FunctionId represents the FunctionId schema from the OpenAPI specification
type FunctionId struct {
	Filepath string `json:"filePath?,omitempty"` // Path to file (Javascript only).
	Functionname string `json:"functionName,omitempty"` // Function name.
	Classname string `json:"className?,omitempty"` // Class name (Java only).
}

// IssueCounts represents the IssueCounts schema from the OpenAPI specification
type IssueCounts struct {
	Results []map[string]interface{} `json:"results"` // A list of issue counts by day
}

// ProjectCountsFilters represents the ProjectCountsFilters schema from the OpenAPI specification
type ProjectCountsFilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// Licensesfilters represents the Licensesfilters schema from the OpenAPI specification
type Licensesfilters struct {
	Filters map[string]interface{} `json:"filters,omitempty"`
}

// Rubygemsrequestpayload represents the Rubygemsrequestpayload schema from the OpenAPI specification
type Rubygemsrequestpayload struct {
	Encoding string `json:"encoding,omitempty"` // the encoding for the manifest files sent.
	Files map[string]interface{} `json:"files"` // The manifest files:
}

// Deletetagbody represents the Deletetagbody schema from the OpenAPI specification
type Deletetagbody struct {
	Value string `json:"value,omitempty"` // Valid tag value.
	Force bool `json:"force,omitempty"` // force delete tag that has entities (default is `false`).
	Key string `json:"key,omitempty"` // Valid tag key.
}

// DepGraphData represents the DepGraphData schema from the OpenAPI specification
type DepGraphData struct {
	Pkgs []interface{} `json:"pkgs"` // Array of package dependencies.
	Schemaversion string `json:"schemaVersion"` // Snyk DepGraph library schema version.
	Graph Graph `json:"graph"`
	Pkgmanager PkgManager `json:"pkgManager"`
}

// MonitorRepository represents the MonitorRepository schema from the OpenAPI specification
type MonitorRepository struct {
	Alias string `json:"alias,omitempty"` // deb, apk and rpm package managers should use an alias to indicate the target Operating System, for example 'debian:10'.
}

// GradleFile represents the GradleFile schema from the OpenAPI specification
type GradleFile struct {
	Contents string `json:"contents"` // The contents of the file, encoded according to the `encoding` field.
}

// ComposerLock represents the ComposerLock schema from the OpenAPI specification
type ComposerLock struct {
	Contents string `json:"contents,omitempty"`
}
