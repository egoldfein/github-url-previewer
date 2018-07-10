# github-url-previewer

# Setup

- add a webhook for the repo in Github and generate a secret token
- generate a personal access token in Github
- generate a key for accessing the Link Preview API: https://my.linkpreview.net/
- create an environment.env file using the environment.env.template file as a template and add the secret token, the personal access token, and the access key for the Link Preview API to their respective spots

# Running locally

- run go build and then ./github-url-previewer to run
- edit a file in the repo and create a PR for the change
- once the PR has been created, the comments will be added automatically 

testing action
