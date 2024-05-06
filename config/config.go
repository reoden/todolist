package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	DbHost     string
	DbName     string
	DbPassword string
	DbPort     string
	DbUser     string
	Charset    string

	EtcdHost string
	EtcdPort string

	UserServiceAddress string
	TaskServiceAddress string
)

func Init() {
	file, err := ini.Load("E:\\golearn\\src\\todolist\\config\\config.ini")
	if err != nil {
		log.Println("Load config file err:", err)
	}

	LoadMySqlData(file)
	LoadEtcd(file)
	LoadServer(file)
}

func LoadServer(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadMySqlData(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	Charset = file.Section("mysql").Key("Charset").String()
}
