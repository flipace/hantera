---
layout: page
title: About
---
hantera is a tool which helps you manage projects which utilize a service
oriented architecture. It's made with passion by [@flipace](https://twitter.com/flipace).

*Tasks which hantera can help you with:*

- setting up a project and all it's dependencies on a machine for development
- versioning and publishing of dependencies as well as whole applications
- deployment of applications including their dependencies at specified versions

### Usage

hantera requires a special hantera.yml (config file) in order to provide you with useful commands for development, versioning or deployment of applications and services based on a service oriented or distributed architecture.

**Example hantera.yml:**
```yaml
version: 0.0.1

name: My Application
description: The best application on the internet.

dependencies:
  react-themeit:
    version: 1.0.0
    description: powerful theming for react components
    repository: git@github.com:flipace/react-themeit.git
  whoport:
    version: 1.0.0
    description: helper tool to easily kill processes which listen on a port
    repository: git@github.com:flipace/whoport.git
  lovlijs:
    version: 1.0.0
    description: lovely, small horizon + react starter kit
    repository: git@github.com:flipace/lovli.js.git
```

### Setup development structure
To quickly setup all repositories on your local machine, simply run
```shell
hantera develop setup
```

This will parse the ```hantera.yml``` in the current working directory, concurrently clone all repositories into ./{repositoryName}/ and afterwards try to install each repositories dependencies (e.g. runs ```yarn``` in repositories which contain a package.json)

If you would like to prevent automatically installing dependencies, you may use the ```--no-deps``` flag.

---

Have questions or suggestions? Feel free to [open an issue on GitHub](https://github.com/flipace/hantera/issues/new) or [ask me on Twitter](https://twitter.com/flipace).

Thanks for reading!
