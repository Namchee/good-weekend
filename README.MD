# Good Weekend

Good Weekend is a simple GitHub Action that automatically rejects any pull requests submitted on weekends. Enjoy your hassle-free weekends!

## Usage

Simply add this actions to your GitHub Workflows job, for example:

```yaml
on:
  pull_request:
    types: [opened] # DO NOT add other triggers!

jobs:
  pr-check:
    runs-on: ubuntu-latest
    steps:
      - name: good-weekend
        uses: Namchee/good-weekend@v{version}
        with:
          access_token: <YOUR_GITHUB_ACCESS_TOKEN_HERE>
```

> Note: **DO NOT** add other triggers other than `pull_request.opened`!

Please refer to [GitHub Workflows Syntax](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions) for more advanced usage.

## Inputs 

This action is customizable with these following inputs (fill it in `with` section in the job steps)

Name | Required? | Default | Description
---- | --------- | ------- | -----------
`access_token` | `true` | | [GitHub's access token](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) which is used to interact with GitHub's API. It is recommended to store this with [GitHub secrets](https://docs.github.com/en/actions/reference/encrypted-secrets)
`timezone` | `false` | `UTC` | Timezone location on [tz database](https://www.iana.org/time-zones). [Cheatsheet](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
`message` | `false` | [See here](./action.yml) | Message to be shown on weekend pull request

For more information, please refer to the [action metadata](./action.yml)

## License

This project is licensed under the [MIT license](./LICENSE)