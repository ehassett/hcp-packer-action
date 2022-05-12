# HCP Packer Action

GitHub Action for automating HCP Packer releases

## Table of Contents
- [HCP Packer Action](#hcp-packer-action)
  - [Table of Contents](#table-of-contents)
  - [Usage](#usage)
    - [Inputs](#inputs)
  - [Author](#author)
  - [License](#license)

## Usage

Add the Action to your GitHub Workflow configuration:
```yaml
name: HCP Packer
on:
  push:

jobs:
  hcp-packer:
    name: hcp-packer
    runs-on: ubuntu-latest
    env:
      HCP_CLIENT_ID: ${{ secrets.HCP_CLIENT_ID }}
      HCP_CLIENT_SECRET: ${{ secrets.HCP_CLIENT_SECRET }}
      HCP_ORGANIZATION_ID: ${{ secrets.HCP_ORGANIZATION_ID }}
      HCP_PROJECT_ID: ${{ secrets.HCP_PROJECT_ID }}

    steps:
      - name: release
        uses: ehassett/hcp-packer-action@main
        with:
          channel: foo
          bucket: bar
```

### Inputs

| Name      | Description                                       | Required | Default |
| --------- | ------------------------------------------------- | -------- | ------- |
| `channel` | release channel to point to most recent iteration | yes      |         |
| `bucket`  | HCP Packer bucket to work with                    | yes      |         |

This Action also relies on the following environment variables:
- HCP_CLIENT_ID
- HCP_CLIENT_SECRET
- HCP_ORGANIZATION_ID
- HCP_PROJECT_ID
See more information about [HCP Service Principals](https://cloud.hashicorp.com/docs/hcp/admin/access-control/service-principals).

## Author

This module is maintained by [Ethan Hassett](https://github.com/ehassett) and the contributors listed on [GitHub](https://github.com/ehassett/hcp-packer-action/graphs/contributors).

## License

Licensed under the [MIT License](LICENSE).
