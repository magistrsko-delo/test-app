package models

import "fmt"

var envStruct *Env

type Env struct {
	key string
}

func (env *Env) SetEnv(key string) {
	(*env).key = key
}

func (env *Env) GetEnv() string {
	return (*env).key
}


func InitEnv(key string)  {
	envStruct = &Env{key: key}
}

func GetEnvStruct() *Env  {
	fmt.Println(envStruct)
	return  envStruct
}