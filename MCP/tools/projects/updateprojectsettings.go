package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/snyk-api/mcp-server/config"
	"github.com/snyk-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdateprojectsettingsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		orgIdVal, ok := args["orgId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: orgId"), nil
		}
		orgId, ok := orgIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: orgId"), nil
		}
		projectIdVal, ok := args["projectId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: projectId"), nil
		}
		projectId, ok := projectIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: projectId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/org/%s/project/%s/settings", cfg.BaseURL, orgId, projectId)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdateprojectsettingsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_org_orgId_project_projectId_settings",
		mcp.WithDescription("Update project settings"),
		mcp.WithString("orgId", mcp.Required(), mcp.Description("Automatically added")),
		mcp.WithString("projectId", mcp.Required(), mcp.Description("Automatically added")),
		mcp.WithBoolean("pullRequestFailOnlyForHighSeverity", mcp.Description("Input parameter: If set to `true`, fail Snyk Test only for high and critical severity vulnerabilities")),
		mcp.WithArray("autoDepUpgradeIgnoredDependencies", mcp.Description("Input parameter: An array of comma-separated strings with names of dependencies you wish Snyk to ignore to upgrade.")),
		mcp.WithString("autoDepUpgradeMinAge", mcp.Description("Input parameter: The age (in days) that an automatic dependency check is valid for")),
		mcp.WithBoolean("pullRequestFailOnAnyVulns", mcp.Description("Input parameter: If set to `true`, fail Snyk Test if the repo has any vulnerabilities. Otherwise, fail only when the PR is adding a vulnerable dependency.")),
		mcp.WithBoolean("pullRequestTestEnabled", mcp.Description("Input parameter: If set to `true`, Snyk Test checks PRs for vulnerabilities.:cq")),
		mcp.WithBoolean("autoDepUpgradeEnabled", mcp.Description("Input parameter: If set to `true`, Snyk will raise dependency upgrade PRs automatically.")),
		mcp.WithObject("autoRemediationPrs", mcp.Description("Input parameter: Defines automatic remediation policies")),
		mcp.WithObject("pullRequestAssignment", mcp.Description("Input parameter: assign Snyk pull requests")),
		mcp.WithString("autoDepUpgradeLimit", mcp.Description("Input parameter: The limit on auto dependency upgrade PRs.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateprojectsettingsHandler(cfg),
	}
}
