package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/snyk-api/mcp-server/config"
	"github.com/snyk-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetlistofissuesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("from=%v", val))
		}
		if val, ok := args["to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["perPage"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("perPage=%v", val))
		}
		if val, ok := args["sortBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortBy=%v", val))
		}
		if val, ok := args["order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order=%v", val))
		}
		if val, ok := args["groupBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("groupBy=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
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
		url := fmt.Sprintf("%s/reporting/issues/%s", cfg.BaseURL, queryString)
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

func CreateGetlistofissuesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_reporting_issues",
		mcp.WithDescription("Get list of issues"),
		mcp.WithString("from", mcp.Required(), mcp.Description("The date you wish to fetch results from, in the format `YYYY-MM-DD`")),
		mcp.WithString("to", mcp.Required(), mcp.Description("The date you wish to fetch results until, in the format `YYYY-MM-DD`")),
		mcp.WithString("page", mcp.Description("The page of results to request")),
		mcp.WithString("perPage", mcp.Description("The number of results to return per page (Maximum: 1000)")),
		mcp.WithString("sortBy", mcp.Description("The key to sort results by")),
		mcp.WithString("order", mcp.Description("The direction to sort results.")),
		mcp.WithString("groupBy", mcp.Description("Set to issue to group the same issue in multiple projects")),
		mcp.WithObject("filters", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetlistofissuesHandler(cfg),
	}
}
