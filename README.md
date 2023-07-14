# Longtea

![image](https://github.com/chungjung-d/longtea/assets/63407866/27745785-fb47-4a4a-8646-116db37e0699)

The longtea is result of container study. It is a simple container system implements by using the Linux namespace system.
Since, longtea is not complete container system, it is not suitable for production use.
Just use it for study.

## Installation

```bash
curl -sSL https://raw.githubusercontent.com/chungjung-d/longtea/develop/install.sh | sudo bash
```

## Command

### Pull

pull the image on the docker hub. If not specified the tag, the latest tag is used.

```bash
$ sudo longtea pull -i <image name>:<tag> # pull the image with tag
$ sudo longtea pull -i <image name>       # pull the image with latest tag
```

**example**

```bash
$ sudo longtea pull -i ubuntu:22.04
```

### Create

Create the container.
If not specified the tag, the latest tag is used and defualt image tag is latest.
The -n option is onptional and if not specified, the container name auto set "**< image name > \_ < tag >**".

```bash
$ sudo longtea create -n <container name> -i <image name>:<tag>  # create the container with container name
$ sudo longtea create -i <image name>:<tag>     # create the container with image name  container name auto set <image name>_<tag>
```

**example**

```bash
$ sudo longtea create -n ubnutu_test -i ubuntu:22.04
```

```bash
$ sudo longtea create -i ubuntu:22.04
```

### Start

Start the created container with container name. If the tag version is not pulled yet, it return error message.

```bash
$ sudo longtea start -n <container name>   # start the container with container name
```

**example**

```bash
$ sudo longtea start -n ubuntu_22.04
```

### Run

Run command is combination of create and start command.
If the tag version is not pulled yet, it return error message.

```bash
$ sudo longtea run -n <container name> -i <image name>:<tag>  # create and start the container with container name
$ sudo longtea run -i <image name>:<tag>     # create and start the container with image name  container name auto set <image name>_<tag>
```

### Remove

Remove the container or image.
Use -i option to remove the image and -c option to remove the container.
If want to remove all image use just image name without tag.

```bash

$ sudo longtea remove -i <image name>:<tag>   # remove the image with tag
$ sudo longtea remove -i <image name>         # remove all image with image name
$ sudo longtea remove -c <container name>     # remove the container
```

### List

List the container or image.
Use -i option to list the image and -c option to list the container or use both -ic option to list all.

```bash
$ sudo longtea list -ic  # list all container and image
$ sudo longtea list -i   # list all image
$ sudo longtea list -c   # list all container
```

## TODO

It need to implement the following features. It will be implemented on version 0.2.0.

- [ ] Implement the network namespace
- [ ] Implement the background run
- [ ] Implement the Squlite DB connection to manage the container
