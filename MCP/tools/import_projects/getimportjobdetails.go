package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snyk-api/mcp-server/config"
	"github.com/snyk-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetimportjobdetailsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		integrationIdVal, ok := args["integrationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: integrationId"), nil
		}
		integrationId, ok := integrationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: integrationId"), nil
		}
		jobIdVal, ok := args["jobId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: jobId"), nil
		}
		jobId, ok := jobIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: jobId"), nil
		}
		url := fmt.Sprintf("%s/org/%s/integrations/%s/import/%s", cfg.BaseURL, orgId, integrationId, jobId)
		req, err := http.NewRequest("GET", url, nil)
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

func CreateGetimportjobdetailsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_org_orgId_integrations_integrationId_import_jobId",
		mcp.WithDescription("Get import job details"),
		mcp.WithString("orgId", mcp.Required(), mcp.Description("The organization ID. The `API_KEY` must have admin access to this organization.")),
		mcp.WithString("integrationId", mcp.Required(), mcp.Description("The unique identifier for the configured integration. This can be found on the [Integration page in the Settings area](https://app.snyk.io/manage/integrations) for all integrations that have been configured.")),
		mcp.WithString("jobId", mcp.Required(), mcp.Description("The ID of the job. This can be found in the Location response header from the corresponding POST request that triggered the import job.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetimportjobdetailsHandler(cfg),
	}
}
