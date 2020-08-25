---
layout: page
title: Installation Guide
permalink: /installation
---

<a name="getting-started"></a>

# Quick Start

Getting Meshery up and running on a locally on Docker-enabled system is easy. Use the Meshery command line interface, `mesheryctl`, to start Meshery on any of its [supported platforms](/docs/installation/platforms).

## Using `mesheryctl`

`mesheryctl` is a command line interface to manage a Meshery deployment. `mesheryctl` allows you to control Meshery's lifecycle with commands like `start`, `stop`, `status`, `reset`. Running `reset` will remove all active container instances, prune pulled images and remove any local volumes created by starting Meshery.

### Mac or Linux

Use your choice of homebrew or bash to install `mesheryctl`. You only need to use one.

#### Homebrew

Install `mesheryctl` and run Meshery on Mac with Homebrew.

**Installing with Homebrew**

To install `mesheryctl`, execute the following commands:

```bash
brew tap layer5io/tap
brew install mesheryctl
mesheryctl system start
```

**Upgrading with Homebrew**

To upgrade `mesheryctl`, execute the following command:

```bash
brew upgrade mesheryctl
```

#### Bash

**Installing with Bash**

Install `mesheryctl` and run Meshery on Mac or Linux with this script:

```bash
curl -L https://git.io/meshery | bash -
```

**Upgrading with Bash**

Upgrade `mesheryctl` and run Meshery on Mac or Linux with this script:

```bash
curl -L https://git.io/meshery | bash -
```

### Windows

#### Installing the `mesheryctl` binary

Download and unzip `mesheryctl` from the [Meshery releases](https://github.com/layer5io/meshery/releases/latest) page. Add `mesheryctl` to your PATH for ease of use. Then, execute:

```bash
./mesheryctl system start
```

#### Scoop

Use [Scoop](https://scoop.sh) to install Meshery on your Windows machine.

**Installing with Scoop**

Add the Meshery Scoop Bucket and install:

```bash
scoop bucket add mesheryctl https://github.com/layer5io/scoop-bucket.git
scoop install mesheryctl
```

**Upgrading with Scoop**

To upgrade `mesheryctl`, execute the following command:

```bash
scoop update mesheryctl
```

---

Upon starting Meshery successfully, instructions to access Meshery will be printed on the sceen.
