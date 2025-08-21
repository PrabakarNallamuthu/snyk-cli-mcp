package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snyk-api/mcp-server/config"
	"github.com/snyk-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListallprojectsnapshotissuepathsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		snapshotIdVal, ok := args["snapshotId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: snapshotId"), nil
		}
		snapshotId, ok := snapshotIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: snapshotId"), nil
		}
		issueIdVal, ok := args["issueId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issueId"), nil
		}
		issueId, ok := issueIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issueId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["perPage"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("perPage=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/org/%s/project/%s/history/%s/issue/%s/paths%s", cfg.BaseURL, orgId, projectId, snapshotId, issueId, queryString)
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

func CreateListallprojectsnapshotissuepathsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_org_orgId_project_projectId_history_snapshotId_issue_issueId_paths",
		mcp.WithDescription("List all project snapshot issue paths"),
		mcp.WithString("orgId", mcp.Required(), mcp.Description("The organization ID. The `API_KEY` must have access to this organization.")),
		mcp.WithString("projectId", mcp.Required(), mcp.Description("The project ID for which to return issue paths.")),
		mcp.WithString("snapshotId", mcp.Required(), mcp.Description("The project snapshot ID for which to return issue paths. If set to `latest`, the most recent snapshot will be used. Use the \"List all project snapshots\" endpoint to find suitable values for this.")),
		mcp.WithString("issueId", mcp.Required(), mcp.Description("The issue ID for which to return issue paths.")),
		mcp.WithString("perPage", mcp.Description("The number of results to return per page (1 - 1000, inclusive).")),
		mcp.WithString("page", mcp.Description("The page of results to return.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListallprojectsnapshotissuepathsHandler(cfg),
	}
}
