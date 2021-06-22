# Postr - The picture posting service

[![CodeQL](https://github.com/DistilledP/postr/actions/workflows/codeql-analysis.yml/badge.svg?branch=master)](https://github.com/DistilledP/postr/actions/workflows/codeql-analysis.yml)

## Introduction

Postr is a small picture posting/hosting system, designed to be lightweight on resources.

It is far from what I would consider a complete project, however, it has served as a useful project to keep my skills up to date.

## Dependencies

The project has been written with a minimal number of dependencies, heck it was written using the Windows Linux compatibility layer.

At a basic level, all you should need is:

- Docker version 19.03+

If you are wanting to run some of the go commands locally, let's face it why wouldn't you?  You will need:

- Go lang v1.16+

## Building the application

Most/all actions can be performed by using the make commands, these are mainly shortcuts to commonly use actions.

To get started, and build the application, the command you will need is:
```bash
make build-app
```

This will start the application build process within a multi-stage docker build.

Once the application build has been completed, you will find the command line uploader executable within the `./bin` at the root of the project.

***Please Note***

The build process has been written in the assumption that the host machine is running a `Linux/amd64` compatible OS.  If you are trying to run the CLI command on a Mac, you will need to recompile the uploader CLI application with the following command:
```bash
make build-cli
```
This will place the executable within the `./bin` directory as before.


## Starting the server

Once you have built the application, the next step is to start the server application.

As the project makes use of `docker-compose` for container orchestration, you simply need to run the command:

```bash
docker-compose up
```

After a couple of seconds the server application should start, to confirm this you should see the following log lines in your console:
```bash
postr_1  | 2021/06/21 21:23:36 GRPC server listening on port 3001
postr_1  | 2021/06/21 21:23:36 HTTP server listening on port 3000
```

This indicates that the application is listening for requests, and ready to roll.

## Usage instructions

As mentioned previously, the binary executable file is located in the `./bin` directory within the application/repository.

To upload a file. you simply need to issue the command:

```bash
./bin/image-upload <path to file>
```

example:
```bash
./bin/image-upload ./test_files/gopher-side_color.png
```

If you wish to save some time and upload a batch of files this can be achieved by adding the additional files to upload at the end of the command - the only requirement is that each image path needs to be separated by white space.

example:
```bash
./bin/image-upload ./test_files/90s_internet.gif ./test_files/gopher-side_color.png ./test_files/lanma256.bmp ./test_files/sad_cat.jpg
```

As the files are uploaded you will receive a summary of which files were successful, for example, the above command resulted in the following output:

```bash
2021/06/21 23:07:36
        Upload        : SUCCESS
        Filename      : 90s_internet.gif
        FileType      : GIF
        Size          : 1640092

2021/06/21 23:07:36
        Upload        : SUCCESS
        Filename      : gopher-side_color.png
        FileType      : PNG
        Size          : 13611

2021/06/21 23:07:36
        Upload        : FAILED
        Filename      : lanma256.bmp
        FileType      : UNKNOWN
        Size          : 0
        Error Message : File format is not accepted, detected: "image/bmp"

2021/06/21 23:07:36
        Upload        : SUCCESS
        Filename      : sad_cat.jpg
        FileType      : JPEG
        Size          : 31850

2021/06/21 23:07:36 Uploaded 3 out of 4 files
```

As you will see, only 3 out of the 4 files were successfully uploaded.  This is because at present the system only supports the following file formats:
- GIF
- JPEG
- PNG

In other words, the most commonly used formats on the internet.

