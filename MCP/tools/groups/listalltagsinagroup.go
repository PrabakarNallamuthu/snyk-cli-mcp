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

func ListalltagsinagroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		groupIdVal, ok := args["groupId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: groupId"), nil
		}
		groupId, ok := groupIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: groupId"), nil
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
		url := fmt.Sprintf("%s/group/%s/tags%s", cfg.BaseURL, groupId, queryString)
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

func CreateListalltagsinagroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_group_groupId_tags",
		mcp.WithDescription("List all tags in a group"),
		mcp.WithString("groupId", mcp.Required(), mcp.Description("The group ID. The `API_KEY` must have access admin to this group.")),
		mcp.WithString("perPage", mcp.Description("The number of results to return (the default is 1000).")),
		mcp.WithString("page", mcp.Description("The offset from which to start returning results from.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListalltagsinagroupHandler(cfg),
	}
}
