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

package service

import (
	"context"
	"log/slog"
	"stalker/internal/client"
)

type GitHubService struct {
	client client.GitHubClient
	logger *slog.Logger
}

func NewGitHubService(client client.GitHubClient, logger *slog.Logger) *GitHubService {
	return &GitHubService{client: client, logger: logger}
}

func (s *GitHubService) GetUsers(ctx context.Context, username, listType string) ([]string, error) {
	s.logger.InfoContext(ctx, "Getting users", "username", username, "listType", listType)

	var allUsers []string
	opt := &client.ListOptions{PerPage: 100}
	for {
		var users []*client.User
		var resp *client.Response
		var err error
		if listType == "followers" {
			users, resp, err = s.client.ListFollowers(ctx, username, opt)
		} else {
			users, resp, err = s.client.ListFollowing(ctx, username, opt)
		}
		if err != nil {
			s.logger.ErrorContext(ctx, "Error getting users", "error", err)
			return nil, err
		}
		for _, user := range users {
			allUsers = append(allUsers, user.GetLogin())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	s.logger.InfoContext(ctx, "Got users", "count", len(allUsers))
	return allUsers, nil
}
