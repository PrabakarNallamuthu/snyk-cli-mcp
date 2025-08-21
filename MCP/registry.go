package main

import (
	"github.com/snyk-api/mcp-server/config"
	"github.com/snyk-api/mcp-server/models"
	tools_dependencies "github.com/snyk-api/mcp-server/tools/dependencies"
	tools_integrations "github.com/snyk-api/mcp-server/tools/integrations"
	tools_users "github.com/snyk-api/mcp-server/tools/users"
	tools_test "github.com/snyk-api/mcp-server/tools/test"
	tools_groups "github.com/snyk-api/mcp-server/tools/groups"
	tools_reporting_api "github.com/snyk-api/mcp-server/tools/reporting_api"
	tools_import_projects "github.com/snyk-api/mcp-server/tools/import_projects"
	tools_organizations "github.com/snyk-api/mcp-server/tools/organizations"
	tools_projects "github.com/snyk-api/mcp-server/tools/projects"
	tools_webhooks "github.com/snyk-api/mcp-server/tools/webhooks"
	tools_audit_logs "github.com/snyk-api/mcp-server/tools/audit_logs"
	tools_entitlements "github.com/snyk-api/mcp-server/tools/entitlements"
	tools_monitor "github.com/snyk-api/mcp-server/tools/monitor"
	tools_licenses "github.com/snyk-api/mcp-server/tools/licenses"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_dependencies.CreateListalldependenciesTool(cfg),
		tools_integrations.CreateUpdateexistingintegrationTool(cfg),
		tools_users.CreateGetorganizationnotificationsettingsTool(cfg),
		tools_users.CreateModifyorganizationnotificationsettingsTool(cfg),
		tools_test.CreateTestdepgraphTool(cfg),
		tools_test.CreateTestcomposer_jsoncomposer_lockfileTool(cfg),
		tools_groups.CreateListalltagsinagroupTool(cfg),
		tools_reporting_api.CreateGettestcountsTool(cfg),
		tools_import_projects.CreateImporttargetsTool(cfg),
		tools_groups.CreateAddamembertoanorganizationwithinagroupTool(cfg),
		tools_test.CreateTestgemfile_lockfileTool(cfg),
		tools_integrations.CreateListTool(cfg),
		tools_integrations.CreateAddnewintegrationTool(cfg),
		tools_organizations.CreateUpdateamembersroleintheorganizationTool(cfg),
		tools_projects.CreateListallaggregatedissuesTool(cfg),
		tools_users.CreateGetprojectnotificationsettingsTool(cfg),
		tools_users.CreateModifyprojectnotificationsettingsTool(cfg),
		tools_integrations.CreateSwitchbetweenbrokertokensTool(cfg),
		tools_organizations.CreateVieworganizationsettingsTool(cfg),
		tools_organizations.CreateUpdateorganizationsettingsTool(cfg),
		tools_integrations.CreateRetrieveTool(cfg),
		tools_integrations.CreateUpdateTool(cfg),
		tools_groups.CreateDeletetagfromgroupTool(cfg),
		tools_test.CreateTestrequirements_txtfileTool(cfg),
		tools_organizations.CreateDeletependinguserprovisionTool(cfg),
		tools_organizations.CreateListpendinguserprovisionsTool(cfg),
		tools_organizations.CreateProvisionausertotheorganizationTool(cfg),
		tools_test.CreateTestmavenfileTool(cfg),
		tools_integrations.CreateProvisionnewbrokertokenTool(cfg),
		tools_test.CreateTestgopkg_tomlgopkg_lockfileTool(cfg),
		tools_webhooks.CreatePingawebhookTool(cfg),
		tools_projects.CreateListallignoresTool(cfg),
		tools_reporting_api.CreateGetlatestprojectcountsTool(cfg),
		tools_projects.CreateDeleteaprojectTool(cfg),
		tools_projects.CreateRetrieveasingleprojectTool(cfg),
		tools_projects.CreateUpdateaprojectTool(cfg),
		tools_audit_logs.CreateGetorganizationlevelauditlogsTool(cfg),
		tools_groups.CreateListallrolesinagroupTool(cfg),
		tools_organizations.CreateListmembersTool(cfg),
		tools_projects.CreateDeleteignoresTool(cfg),
		tools_projects.CreateRetrieveignoreTool(cfg),
		tools_projects.CreateAddignoreTool(cfg),
		tools_projects.CreateReplaceignoresTool(cfg),
		tools_groups.CreateListallorganizationsinagroupTool(cfg),
		tools_reporting_api.CreateGetlatestissuecountsTool(cfg),
		tools_organizations.CreateSetnotificationsettingsTool(cfg),
		tools_organizations.CreateGet_org_orgid_notification_settingsTool(cfg),
		tools_projects.CreateListallprojectsTool(cfg),
		tools_audit_logs.CreateGetgrouplevelauditlogsTool(cfg),
		tools_integrations.CreateCloneanintegrationwithsettingsandcredentialsTool(cfg),
		tools_projects.CreateListallprojectsnapshotaggregatedissuesTool(cfg),
		tools_projects.CreateListallprojectsnapshotsTool(cfg),
		tools_test.CreateTestvendor_jsonfileTool(cfg),
		tools_entitlements.CreateListallentitlementsTool(cfg),
		tools_projects.CreateGetprojectdependencygraphTool(cfg),
		tools_reporting_api.CreateGetprojectcountsTool(cfg),
		tools_integrations.CreateDeletecredentialsTool(cfg),
		tools_reporting_api.CreateGetlistofissuesTool(cfg),
		tools_test.CreateGet_test_sbt_groupid_artifactid_versionTool(cfg),
		tools_organizations.CreateListalltheorganizationsauserbelongstoTool(cfg),
		tools_test.CreateTestpackage_jsonyarn_lockfileTool(cfg),
		tools_integrations.CreateGetexistingintegrationbytypeTool(cfg),
		tools_projects.CreateDeactivateTool(cfg),
		tools_import_projects.CreateGetimportjobdetailsTool(cfg),
		tools_reporting_api.CreateGetlistoflatestissuesTool(cfg),
		tools_monitor.CreateMonitordepgraphTool(cfg),
		tools_projects.CreateListallprojectissuepathsTool(cfg),
		tools_organizations.CreateRemoveorganizationTool(cfg),
		tools_organizations.CreateInviteusersTool(cfg),
		tools_projects.CreateCreatejiraissueTool(cfg),
		tools_entitlements.CreateGetanorganizationsentitlementvalueTool(cfg),
		tools_test.CreateTestforissuesinapublicpackagebygroupnameandversionTool(cfg),
		tools_test.CreateTestforissuesinapublicgembynameandversionTool(cfg),
		tools_webhooks.CreateDeleteawebhookTool(cfg),
		tools_webhooks.CreateRetrieveawebhookTool(cfg),
		tools_projects.CreateListallprojectsnapshotissuepathsTool(cfg),
		tools_projects.CreateMoveprojecttoadifferentorganizationTool(cfg),
		tools_test.CreateTestforissuesinapublicpackagebynameandversionTool(cfg),
		tools_test.CreateTestsbtfileTool(cfg),
		tools_users.CreateGetmydetailsTool(cfg),
		tools_licenses.CreateListalllicensesTool(cfg),
		tools_organizations.CreateRemoveamemberfromtheorganizationTool(cfg),
		tools_organizations.CreateUpdateamemberintheorganizationTool(cfg),
		tools_projects.CreateDeleteprojectsettingsTool(cfg),
		tools_projects.CreateListprojectsettingsTool(cfg),
		tools_projects.CreateUpdateprojectsettingsTool(cfg),
		tools_test.CreateTestforissuesinapublicpackagebygroupidartifactidandversionTool(cfg),
		tools_test.CreateGet_test_pip_packagename_versionTool(cfg),
		tools_projects.CreateActivateTool(cfg),
		tools_groups.CreateListallmembersinagroupTool(cfg),
		tools_projects.CreateApplyingattributesTool(cfg),
		tools_projects.CreateListalljiraissuesTool(cfg),
		tools_groups.CreateUpdategroupsettingsTool(cfg),
		tools_groups.CreateViewgroupsettingsTool(cfg),
		tools_users.CreateGetuserdetailsTool(cfg),
		tools_projects.CreateAddatagtoaprojectTool(cfg),
		tools_test.CreateTestpackage_jsonpackage_lock_jsonfileTool(cfg),
		tools_reporting_api.CreateGetissuecountsTool(cfg),
		tools_test.CreateTestgradlefileTool(cfg),
		tools_webhooks.CreateListwebhooksTool(cfg),
		tools_webhooks.CreateCreateawebhookTool(cfg),
		tools_organizations.CreateCreateaneworganizationTool(cfg),
		tools_projects.CreateRemoveatagfromaprojectTool(cfg),
	}
}
