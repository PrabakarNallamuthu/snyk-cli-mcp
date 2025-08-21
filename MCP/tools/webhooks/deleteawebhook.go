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

func DeleteawebhookHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		webhookIdVal, ok := args["webhookId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: webhookId"), nil
		}
		webhookId, ok := webhookIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: webhookId"), nil
		}
		url := fmt.Sprintf("%s/org/%s/webhooks/%s", cfg.BaseURL, orgId, webhookId)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDeleteawebhookTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_org_orgId_webhooks_webhookId",
		mcp.WithDescription("Delete a webhook"),
		mcp.WithString("orgId", mcp.Required(), mcp.Description("Automatically added")),
		mcp.WithString("webhookId", mcp.Required(), mcp.Description("Automatically added")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteawebhookHandler(cfg),
	}
}
