# Markdown GitHub Stars Updater

This Go program reads a Markdown file containing GitHub repository links and updates the links with information about star counts. 

#### Example
The initial part of the markdown file:
>- [APITree](https://www.apitree.com/) - A tool for managing and sharing API specifications, with version control, API testing automation, and integration with popular API tools.
>- [DapperDox](https://github.com/DapperDox/dapperdox) - An open-source API documentation generator and server for OpenAPI/Swagger specifications, with customizable documentation, automated updates, and easy sharing.
>- [OpenAPI Explorer](https://github.com/Rhosys/openapi-explorer) - A tool for generating user interfaces from OpenAPI specifications, making it easier for software engineers to visualize and interact with APIs.
>- [RapiDoc](https://github.com/rapi-doc/RapiDoc) - A tool that generates customizable, interactive API documentation from OpenAPI Specification, with a range of design options.
>- [Redoc](https://github.com/Redocly/redoc) - An open-source tool for generating documentation from OpenAPI (fka Swagger) definitions, with customizable themes, language support, and branding.

After processing:
>- [APITree](https://www.apitree.com/) - A tool for managing and sharing API specifications, with version control, API testing automation, and integration with popular API tools.
>- [DapperDox (⭐377)](https://github.com/DapperDox/dapperdox) - An open-source API documentation generator and server for OpenAPI/Swagger specifications, with customizable documentation, automated updates, and easy sharing.
>- [OpenAPI Explorer (⭐213)](https://github.com/Rhosys/openapi-explorer) - A tool for generating user interfaces from OpenAPI specifications, making it easier for software engineers to visualize and interact with APIs.
>- [RapiDoc (⭐1.3k)](https://github.com/rapi-doc/RapiDoc) - A tool that generates customizable, interactive API documentation from OpenAPI Specification, with a range of design options.
>- [Redoc (⭐20k)](https://github.com/Redocly/redoc) - An open-source tool for generating documentation from OpenAPI (fka Swagger) definitions, with customizable themes, language support, and branding.

## Usage
The program will update the label links in Markdown file with star counts for each repository link.
#### Build from sources

1. Clone the repository:
```sh
git clone https://github.com/stn1slv/markdown-github-stars-updater.git
cd markdown-github-stars-updater
```

2. Build and run the program:
```sh
go build
./markdown-github-stars-updater path/to/your/markdown/file.md
 ```
Replace path/to/your/markdown/file.md with the path to your Markdown file.

#### Download compiled

The last compiled version is available in [the releases section](https://github.com/stn1slv/markdown-github-stars-updater/releases/latest).

#### GitHub Actions pipeline

Please see [the pipeline example](https://github.com/stn1slv/awesome-integration/blob/main/.github/workflows/github-stars.yml) for my awesome-integration project.

## Description
This program utilizes the GitHub API to fetch star counts for GitHub repositories and updates the links in the provided Markdown file. It supports different star count formats based on the number of stars:

- If the star count is less than 1000, the exact star count is shown (e.g., "⭐35").
- If the star count is between 1000 and 9999, it is displayed in the format "1.1k" for 1100 stars.
- If the star count is 10000 or more, it is displayed in the format "10k" for 10000 stars.
## Requirements
- Go programming language (https://golang.org/dl/)
## Configuration
To use this program, you need to provide your GitHub access token for API requests. Set the GITHUB_TOKEN environment variable before running the program.

## License
This project is licensed under the MIT License. See LICENSE for more information.

## Acknowledgements
Special thanks to the GitHub community and open-source contributors.
