# Longtea

## Description

The longtea is result of container study. It is a simple container system implements by using the Linux namespace system.
Since, longtea is not complete container system, it is not suitable for production use.
Just use it for study.

## Command

### Pull

pull the image on the docker hub.

```bash
$ sudo longtea pull - i <image name>
```

**example**

```bash
$ sudo longtea pull - i ubuntu
```

### Create 
Create the container.

```bash
$ sudo longtea create - n <container name> - i <image name>
```

**example**

```bash
$ sudo longtea create - n ubnutu_test - i ubuntu
```

### Start
Start the created container.

```bash
$ sudo longtea start - n <container name>
```

**example**

```bash
$ sudo longtea start - n ubnutu_test
```