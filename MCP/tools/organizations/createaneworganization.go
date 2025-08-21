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

func CreateaneworganizationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		url := fmt.Sprintf("%s/org", cfg.BaseURL)
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

func CreateCreateaneworganizationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_org",
		mcp.WithDescription("Create a new organization"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the new organization")),
		mcp.WithString("sourceOrgId", mcp.Description("Input parameter: The id of an organization to copy settings from.\n\nIf provided, this organization must be associated with the same group.\n\nThe items that will be copied are: \nSource control integrations (GitHub, GitLab, BitBucket)\n\\+ Container registries integrations (ACR, Docker Hub, ECR, GCR)\n\\+ Container orchestrators integrations (Kubernetes)\n\\+ PaaS and Serverless Integrations (Heroku, AWS Lambda)\n\\+ Notification integrations (Slack, Jira)\n\\+ Policies\n\\+ Ignore settings\n\\+ Language settings\n\\+ Infrastructure as Code settings\n\\+ Snyk Code settings\n\nThe following will not be copied across:\nService accounts\n\\+ Members\n\\+ Projects\n\\+ Notification preferences")),
		mcp.WithString("groupId", mcp.Description("Input parameter: The group ID. The `API_KEY` must have access to this group.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateaneworganizationHandler(cfg),
	}
}
