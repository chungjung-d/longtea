package config

type ctxKey string

var ContainerDirPath ctxKey = "containerDirPath"
var ImageDirPath ctxKey = "imageDirPath"

var Container string = "containers"
var Image string = "images"

var RemoveContainerName ctxKey = "removeContainerName"
var RemoveImageName ctxKey = "removeImageName"

var CreateContainerName ctxKey = "createContainerName"
var CreateContainerOriginImageName ctxKey = "createContainerOriginImageName"
