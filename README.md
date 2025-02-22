# go-checkout-script
A github action to clone a repository without using node/typescript as required in [actions/checkout](https://github.com/actions/checkout). This repo is written in Golang but supplies builds for linux and windows (amd64) to prevent having to pull dependencies in pipelines.

# Issues
This is still a very new repo in active development and may have breaking changes until v1. At the moment only public repo's can be pulled and pulling specific branches or tags is not fully implemented.

# Binary sizes
Here are the current sizes of the built binaries
| Version (OS) | Size (mb) |
| - | - | 
| linux | 11.8 | 
| windows | 12.2 |

# Usage
basic usage to clone the repository that triggered the event, the repo input can be supplied to pull other repositories
```
uses: jake-young-dev/go-checkout-script@master

uses: jake-young-dev/go-checkout-script@master
with:
    repo: https://github.com/jake-young-dev/go-install-script
```