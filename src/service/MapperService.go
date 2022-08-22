package service

import (
	"Themis/src/config"
	"Themis/src/entity"
	"Themis/src/mapper"
	"time"
)

func LoadDatabase() (E any) {
	defer func() {
		E = recover()
	}()
	serverModels, deleteServerModels := mapper.SelectAllServers()
	for i := 0; i < len(serverModels); i++ {
		model := &entity.ServerModel{
			IP:        serverModels[i].IP,
			Port:      serverModels[i].Port,
			Name:      serverModels[i].Name,
			Time:      serverModels[i].Time,
			Colony:    serverModels[i].Colony,
			Namespace: serverModels[i].Namespace,
		}
		RegisterServer(model)
	}
	for i := 0; i < len(deleteServerModels)-1; i++ {
		model := &entity.ServerModel{
			IP:        serverModels[i].IP,
			Port:      serverModels[i].Port,
			Name:      serverModels[i].Name,
			Time:      serverModels[i].Time,
			Colony:    serverModels[i].Colony,
			Namespace: serverModels[i].Namespace,
		}
		DeleteServer(model)
	}
	return nil
}

func Persistence() (E any) {
	defer func() {
		E = recover()
	}()
	for {
		time.Sleep(time.Duration(config.PersistenceTime) * time.Second)
		mapper.DeleteAllServer()
		mapper.StorageList(InstanceList, mapper.NORMAL)
		mapper.StorageList(DeleteInstanceList, mapper.DELETE)
	}
}

func DeleteMapper(model *entity.ServerModel) (E any) {
	defer func() {
		E = recover()
	}()
	mapper.DeleteServer(model)
	return nil
}
