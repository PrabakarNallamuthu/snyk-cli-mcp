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

func AddignoreHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/org/%s/project/%s/ignore/%s", cfg.BaseURL, orgId, projectId, issueId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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

func CreateAddignoreTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_org_orgId_project_projectId_ignore_issueId",
		mcp.WithDescription("Add ignore"),
		mcp.WithString("orgId", mcp.Required(), mcp.Description("Automatically added")),
		mcp.WithString("projectId", mcp.Required(), mcp.Description("Automatically added")),
		mcp.WithString("issueId", mcp.Required(), mcp.Description("Automatically added")),
		mcp.WithBoolean("disregardIfFixable", mcp.Required(), mcp.Description("Input parameter: Only ignore the issue if no upgrade or patch is available.")),
		mcp.WithString("expires", mcp.Description("Input parameter: The timestamp that the issue will no longer be ignored.")),
		mcp.WithString("ignorePath", mcp.Description("Input parameter: The path to ignore (default is `*` which represents all paths).")),
		mcp.WithString("reason", mcp.Description("Input parameter: The reason that the issue was ignored.")),
		mcp.WithString("reasonType", mcp.Required(), mcp.Description("Input parameter: The classification of the ignore.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AddignoreHandler(cfg),
	}
}
