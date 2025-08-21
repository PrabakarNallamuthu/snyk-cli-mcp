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

func TestforissuesinapublicpackagebygroupidartifactidandversionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		artifactIdVal, ok := args["artifactId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: artifactId"), nil
		}
		artifactId, ok := artifactIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: artifactId"), nil
		}
		versionVal, ok := args["version"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: version"), nil
		}
		version, ok := versionVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: version"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["org"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("org=%v", val))
		}
		if val, ok := args["repository"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("repository=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/test/maven/%s/%s/%s%s", cfg.BaseURL, groupId, artifactId, version, queryString)
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

func CreateTestforissuesinapublicpackagebygroupidartifactidandversionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_test_maven_groupId_artifactId_version",
		mcp.WithDescription("Test for issues in a public package by group id, artifact id and version"),
		mcp.WithString("groupId", mcp.Required(), mcp.Description("The package's group ID.")),
		mcp.WithString("artifactId", mcp.Required(), mcp.Description("The package's artifact ID.")),
		mcp.WithString("version", mcp.Required(), mcp.Description("The package version to test.")),
		mcp.WithString("org", mcp.Description("The organization to test the package with. See \"The Snyk organization for a request\" above.")),
		mcp.WithString("repository", mcp.Description("The Maven repository hosting this package. The default value is Maven Central. More than one value is supported, in order.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TestforissuesinapublicpackagebygroupidartifactidandversionHandler(cfg),
	}
}
