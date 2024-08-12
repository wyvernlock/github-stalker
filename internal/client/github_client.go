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

package client

import (
	"context"
	"log/slog"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

type User = github.User
type Response = github.Response
type ListOptions = github.ListOptions

type GitHubClient interface {
	ListFollowers(ctx context.Context, username string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
	ListFollowing(ctx context.Context, username string, opts *github.ListOptions) ([]*github.User, *github.Response, error)
}

type githubClient struct {
	client *github.Client
	logger *slog.Logger
}

func NewGitHubClient(ctx context.Context, token string, logger *slog.Logger) GitHubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &githubClient{client: client, logger: logger}
}

func (c *githubClient) ListFollowers(ctx context.Context, username string, opts *github.ListOptions) ([]*github.User, *github.Response, error) {
	c.logger.InfoContext(ctx, "Listing followers", "username", username)

	users, resp, err := c.client.Users.ListFollowers(ctx, username, opts)
	if err != nil {
		c.logger.ErrorContext(ctx, "Error listing followers", "username", username, "error", err)
		return nil, nil, err
	}

	c.logger.InfoContext(ctx, "Successfully listed followers", "username", username, "count", len(users))
	return users, resp, nil
}

func (c *githubClient) ListFollowing(ctx context.Context, username string, opts *github.ListOptions) ([]*github.User, *github.Response, error) {
	c.logger.InfoContext(ctx, "Listing following", "username", username)

	users, resp, err := c.client.Users.ListFollowing(ctx, username, opts)
	if err != nil {
		c.logger.ErrorContext(ctx, "Error listing following", "username", username, "error", err)
		return nil, nil, err
	}

	c.logger.InfoContext(ctx, "Successfully listed following", "username", username, "count", len(users))
	return users, resp, nil
}
