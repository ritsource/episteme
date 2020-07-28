# What is it
The idea is to build a platform for sharing what you learn.

# What is it right now
Just a very simple prototype (kind of a blog) where I'm gonna be shaing my learnings (as little curations).

The "curation" part is one of the core parts of what me and my friends were initially trying to build. The idea was kind of like a github for general learning. Hopefully, we will be able to put together something simillar sooner or later.
<!--- More about that here ... -->

# How you too can share your learnings
Ofcourse there's the everold option of building & hosting all your data on a website. That's great :). But, feel free to use the prototype I have built. Here's a step by step process to get up and running with the codebase (FYI, you don't have to write code). 

Here's a screenshot of the application ...

![screenshot](https://raw.githubusercontent.com/ritsource/episteme/master/prototype/assets/screenshots/screenshot.0.png)

Link here ...

# Getting started with the codebase

> Please don't ask why did I use protocol buffers for this project. I was building schema for a richer application. Now it's stuck with us :( :p.

This application takes a `.yaml` as input where you define all teh link and stuff you want to share. Here's an example [configuraton.yml]() file for your reference.

### Requirements
1. [Golang](https://golang.org)
2. [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation)

### Clone the repository
```
git clone git@github.com:ritsource/episteme.git
```

### Config File
Modify `config.yaml` as per your need. 

### Generate protobuf code (go)
> You can skip this step.

```shell
bash prototype/scripts/gen-protobuf.sh
```

### Store protobuf-messages on a destination file

```shell
bash prototype/scripts/decode_config_file.sh
```

### Start server
To start server in development mode
```shell
bash prototype/scripts/run-server-meta.sh

# or using tusk (https://github.com/rliebz/tusk)
cd prototype
tusk dev
```







