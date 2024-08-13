# GithubStalker

GithubStalker is a Go-based CLI tool that helps you identify GitHub users who you follow but aren't following you back.

## Features

* Retrieves your followers and the users you're following from GitHub using the GitHub API.
* Compares the two lists to identify users who aren't reciprocating your follow.
* Presents a clear list directly in your terminal.

## Prerequisites

* Go (installed and configured)
* A GitHub personal access token with the `user:read:follows` scope.

## Setup

1. Clone this repository.
2. Duplicate the .env.example file located at /internal/.env-example and rename it to .env in the same directory. Replace GITHUB_TOKEN with your personal access token and GITHUB_USERNAME with your GitHub username in the .env file. 
3. Run `go mod tidy` to install dependencies.

## Usage

Simply run the application from the project root:

```
go run main.go
```
The output will display a list of GitHub usernames who are not following you back.

## Notes
>The application uses the slog package for structured logging. </br> Error handling is implemented to provide feedback in case of issues. </br> Make sure to keep your GitHub personal access token secure.

## Contributing

Contributions are welcome! If you notice any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request. Your help in fixing bugs, adding features, or enhancing documentation is greatly appreciated. Thanks for contributing and helping make this project better!
