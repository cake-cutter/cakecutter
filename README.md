<div align="center">
<!-- logo -->
<img src = "https://github.com/Dhravya/Cakecutter/raw/main/images/cakecutter.png" width="300">
<h1 align="center">Cakecutter</h1>
<img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" /><br>
Create projects from pre-built cakes (templates) in seconds!
</div>

***

[Read the full Documentation](https://docs.cakes.run)

## 👀 What is Cakecutter?
Sometimes, the most difficult thing is to just get started with a project. Cakecutter is a tool that helps you to cut the cake and start your amazing project instantly. 

What Cakecutter does:
- Users can [publish](http://docs.cakes.run/6-publish/), [create](http://docs.cakes.run/3-create/) or [use a cake](http://docs.cakes.run/2-usage/) from [Cakes.run](https://cakes.run). Cakes are basically TOML files which contain all the information needed to create a project. 
- According to the information in the `Cakefile`, Cakecutter will create all the files and (you can also fill them with content) in the correct location.
- Setup commands (installing dependencies, etc) can be defined in the `Cakefile`. These commands are run after the files are generated.
- Cakecutter can ask questions to the user and take input. The input can then be used as variables for the project template. [Read the docs here](http://docs.cakes.run/4-advance/)
- Read the [Basic example cake](http://docs.cakes.run/5-example/) and [/examples/Python.toml] the Python Cake to see how Cakes are written.

## Installation

### Using `go`
The cli is written in `go`, so if you have `go` installed run
```
go install github.com/cake-cutter/cc@latest
```

### Using `npm`

If you have `npm` you can install `cakecutter` by running

<details>
  <summary>For Windows</summary>

```
npm install -g cakecutter
```

</details>

<details>
  <summary>For MacOS</summary>

```
npm install -g cc-for-mac
```

</details>

<details>
  <summary>For Linux</summary>

```
npm install -g cc-for-linux
```

</details>

> Testing hasn't been done for `macOS` and `linux` yet. Please [create an issue](/issues) if you find any bug.

### Using binary executables

Download the binary executables from `bin/<your_os>`.

## Usage
```
cc help
```

## How to find cakes

You can find cakes for your use case on [Cakes.run](https://cakes.run). You can also publish cakes yourself. 

## Read the docs!
Read the [Docs](http://docs.cakes.run) to learn how to use the cli, create and publish cakes yourself.

Find cakes made by other users on [Cakes.run](https://cakes.run)


## Contributing
Cakecutter is *fully* open source. If you want to contribute, please create an issue on the appropriate [repository](https://github.com/cake-cutter) and we will assign the task to someone (or you).

## Help and community
If you need any help, or want to ask questions, or suggest features, please head over to the [Discussions tab](https://github.com/cake-cutter/cc/discussions)

You can also join our [Discord server](https://discord.gg/z7MZYhmx6w) where we have a community of developers, ready to help you out.