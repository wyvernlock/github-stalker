/*
GithubStalker
Go-based CLI tool that helps you identify GitHub users who you follow but aren't following you back.

Copyright (C) 2024 Rian

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*
You can use your GitHub token either directly in the source code or through environment variables. Hardcoding the token in the code can expose it to anyone with access to the codebase, which is a significant security risk and should be avoided in production environments.

Using environment variables is a more secure practice, as it keeps sensitive information separate from the source code and reduces the risk of accidental exposure. However, ensure that environment variables are managed properly to avoid unintentional disclosure through logs or misconfiguration.

For enhanced security, prefer using environment variables for managing tokens and follow best practices for handling sensitive data.
*/

package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"stalker/internal/client"
	"stalker/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../internal/.env")
	if err != nil {
		slog.Error("Error loading .env file", "err", err)
		os.Exit(1)
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		slog.Error("GITHUB_TOKEN environment variable is required")
		os.Exit(1)
	}

	username := os.Getenv("GITHUB_USERNAME")
	if username == "" {
		slog.Error("GITHUB_USERNAME environment variable is required")
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	ctx := context.Background()

	logger.Info("Initializing GitHub client")
	githubClient := client.NewGitHubClient(ctx, token, logger)

	logger.Info("Creating GitHub service")
	service := service.NewGitHubService(githubClient, logger)

	logger.Info("Fetching followers", "username", username)
	followers, err := service.GetUsers(ctx, username, "followers")
	if err != nil {
		logger.Error("Error getting followers", "error", err)
		os.Exit(1)
	}

	logger.Info("Fetching following", "username", username)
	following, err := service.GetUsers(ctx, username, "following")
	if err != nil {
		logger.Error("Error getting following", "error", err)
		os.Exit(1)
	}

	logger.Info("Comparing users")
	notFollowingBack := client.CompareUsers(following, followers)

	logger.Info("Printing results")
	fmt.Println("Users not following you back:")
	for _, user := range notFollowingBack {
		fmt.Println(user)
		logger.Debug("User not following back", "username", user)
	}

	logger.Info("Process completed successfully")
}
