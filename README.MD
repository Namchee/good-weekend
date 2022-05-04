# Good Weekend

[![Go Version](https://img.shields.io/github/go-mod/go-version/namchee/good-weekend)](https://github.com/Namchee/good-weekend)

Good Weekend is a simple GitHub Action that automatically rejects any pull requests submitted on weekends. Enjoy your hassle-free weekends!

## Usage

Simply add this actions to your GitHub Workflows job, for example:

```yaml
on:
  pull_request:

jobs:
  pr-check:
    runs-on: ubuntu-latest
    steps:
      - name: good-weekend
        uses: Namchee/good-weekend@v{version}
        with:
          access_token: <YOUR_GITHUB_ACCESS_TOKEN_HERE>
```

Please refer to [GitHub Workflows Syntax](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions) for more advanced usage.

## Inputs 

This action is customizable with these following inputs (fill it in `with` section in the job steps)

Name | Required? | Default | Description
---- | --------- | ------- | -----------
`access_token` | `true` | | [GitHub's access token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) which is used to interact with GitHub's API. It is recommended to store this with [GitHub secrets](https://docs.github.com/en/actions/reference/encrypted-secrets)
`timezone` | `false` | `UTC` | Timezone location on [tz database](https://www.iana.org/time-zones). [Cheatsheet](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
`message` | `false` | [See here](./action.yml) | Message to be shown on weekend-submitted pull requests
`label` | `false` | `good-weekend` | Label to be added on weekend-submitted pull requests

For more information, please refer to the [action metadata](./action.yml)

## Forked Repository

By default, the `pull_request` event does not allow GitHub actions to be executed from a forked repository due to [GitHub repository access exploit via GitHub Action](https://securitylab.github.com/research/github-actions-preventing-pwn-requests/). However, this can be circumenvented changing the [event target](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows) from `pull_request` to `pull_request_target` which changes the execution context from the fork to the base repository. Below is the example of action configuration using `pull_request_target`.

```yaml
on:
  pull_request_target:

jobs:
  pr-check:
    runs-on: ubuntu-latest
    steps:
      - name: good-weekend
        uses: Namchee/good-weekend@v{version}
        with:
          access_token: <YOUR_GITHUB_ACCESS_TOKEN_HERE>
```

## License

This project is licensed under the [MIT license](./LICENSE)