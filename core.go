package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"log"
)

// KillProcessByPort 根据端口号关闭进程
func KillProcessByPort(port uint32) {
	entity := GetConnectionByPort(port)
	if entity == nil {
		return
	}
	pr, err := process.NewProcess(entity.Pid)
	if err != nil {
		log.Fatal(err)
	}
	err = pr.Terminate()
	if err != nil {
		log.Fatal(err)
	}
}

// GetConnectionByPort 根据端口号获取连接
func GetConnectionByPort(port uint32) *ConnectionEntity {
	connections := GetConnections()
	for _, con := range connections {
		if con.Port == port {
			return &con
		}
	}
	return nil
}

// GetConnections 获取处于LISTEN状态的TCP连接
func GetConnections() []ConnectionEntity {
	connections, _ := net.Connections("tcp")
	entities := make([]ConnectionEntity, 0)
	for _, con := range connections {
		if con.Status == "LISTEN" {
			entities = append(entities, ConnectionEntity{
				Port: con.Laddr.Port,
				Pid:  con.Pid,
				Name: GetPidName(con.Pid),
			})
		}
	}
	return entities
}

func List(port uint32) {
	connections := GetConnections()
	if port == 0 {
		fmt.Printf("%-10s%-10s%s\n", "port", "pid", "name")
		for _, con := range connections {
			fmt.Printf("%-10d%-10d%s\n", con.Port, con.Pid, con.Name)
		}
	} else {
		for _, con := range connections {
			if con.Port == port {
				fmt.Printf("%-10s%-10s%s\n", "port", "pid", "name")
				fmt.Printf("%-10d%-10d%s\n", con.Port, con.Pid, con.Name)
				break
			}
		}
	}
}

// GetPidName 根据PID获取进程名称
func GetPidName(pid int32) string {
	pro, err := process.NewProcess(pid)
	if err != nil {
		log.Fatal(err)
	}
	var name string
	if name, err = pro.Name(); err != nil {
		log.Fatal(err)
	}
	return name
}

type ConnectionEntity struct {
	Port uint32 `json:"port"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
}
