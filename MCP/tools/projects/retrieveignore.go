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

func RetrieveignoreHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		issueIdVal, ok := args["issueId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issueId"), nil
		}
		issueId, ok := issueIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issueId"), nil
		}
		url := fmt.Sprintf("%s/org/%s/project/%s/ignore/%s", cfg.BaseURL, orgId, projectId, issueId)
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

func CreateRetrieveignoreTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_org_orgId_project_projectId_ignore_issueId",
		mcp.WithDescription("Retrieve ignore"),
		mcp.WithString("orgId", mcp.Required(), mcp.Description("The organization ID to modify ignores for. The `API_KEY` must have access to this organization.")),
		mcp.WithString("projectId", mcp.Required(), mcp.Description("The project ID to modify ignores for.")),
		mcp.WithString("issueId", mcp.Required(), mcp.Description("The issue ID to modify ignores for. Can be a vulnerability or a license Issue.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RetrieveignoreHandler(cfg),
	}
}
