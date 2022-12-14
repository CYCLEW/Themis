package controller

import (
	"Themis/src/config"
	"Themis/src/controller/util"
	"Themis/src/entity"
	"Themis/src/exception"
	"Themis/src/service"
	"Themis/src/util/encryption"
	"Themis/src/util/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GerNamespacesController
// @Summary     获取全部命名空间
// @Description 由管理者调用的获取当前所有服务信息。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200   {object} entity.ResultModel{data=[]string} "返回命名空间列表名称"
// @Router      /operator/CURD/getNamespaces [GET]
func GerNamespacesController(c *gin.Context) {
	namespaces, err := service.GetNamespaces()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(namespaces))
}

// GetColoniesController
// @Summary     获取指定命名空间下的集群
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200   {object} entity.ResultModel{data=[]string} "返回集群列表名称"
// @Router      /operator/CURD/getColonies [POST]
func GetColoniesController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("GetColoniesController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	colony, err := service.GetColonyByNamespace(Server.Namespace)
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(colony))
}

// GetColoniesAndServerController
// @Summary     获取指定命名空间下的集群和服务名
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200   {object} entity.ResultModel{data=map[string][]string} "返回集群服务列表名称"
// @Router      /operator/CURD/getColoniesInstances [POST]
func GetColoniesAndServerController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("GetColoniesAndServerController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	colony, err := service.GetColonyAndServerByNamespace(Server.Namespace)
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(colony))
}

// GetAllServersController
// @Summary     获取全部服务实例
// @Description 由管理者调用的获取当前所有服务信息。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200   {object} entity.ResultModel{} "返回服务全体封装"
// @Router      /operator/CURD/getInstances [GET]
func GetAllServersController(c *gin.Context) {
	servers, err := service.GetInstances()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(servers))
}

// GetInstancesByConditionListController
// @Summary     获取指定条件下的服务实例服务实例--模糊查询
// @Description 获取命名空间与集群条件下的服务实例。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Param       Model body     entity.ServerModel                            true "封装的条件参数"
// @Success     200 {object} entity.ResultModel{data=[]entity.ServerModel} "返回服务实例切片数组"
// @Router      /operator/CURD/getInstancesByCondition [POST]
func GetInstancesByConditionListController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("GetInstancesByConditionListController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	list, err := service.GetInstancesByNamespaceAndColony(Server)
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(list))
}

// GetInstancesByConditionMapController
// @Summary     获取指定条件下的服务实例服务实例--精确查询
// @Description 获取命名空间与集群条件下的服务实例。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Param       Model body     entity.ServerModel                            true "封装的条件参数"
// @Success     200 {object} entity.ResultModel{data=map[string][]entity.ServerModel} "返回服务实例map切片数组"
// @Router      /operator/CURD/getInstancesByConditionMap [POST]
func GetInstancesByConditionMapController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("GetInstancesByConditionMapController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	if Server.Namespace == "" || Server.Colony == "" {
		exception.HandleException(exception.NewUserError("GetInstancesByConditionMapController", "参数绑定错误-命名空间或集群为空"))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-命名空间或集群为空"))
		return
	}
	list, err := service.GetInstanceByNamespaceAndColony(Server)
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(list))
}

// DeleteInstanceController
// @Summary     删除服务实例并拉入黑名单
// @Description 由管理者调用删除服务实例并拉入黑名单。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Param       Model body     entity.ServerModel true "被删除的服务实例信息"
// @Success     200   {object} entity.ResultModel "返回true或false"
// @Router      /operator/CURD/deleteInstance [delete]
func DeleteInstanceController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("DeleteInstanceController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	Assert1, err1 := service.CheckServer(Server)
	if err1 != nil {
		util.Handle(err1, c)
		return
	}
	if Assert1 {
		Assert2, err2 := service.DeleteServer(Server)
		if err2 != nil {
			util.Handle(err2, c)
			return
		}
		if Assert2 {
			c.JSON(http.StatusOK, entity.NewSuccessResult(Assert2))
		}
	} else {
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "实例不存在"))
	}
}

// DeleteInstanceByColonyController
// @Summary     删除地区集群实例并拉入黑名单
// @Description 由管理者调用删除地区集群实例并拉入黑名单。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Param       Model body     entity.ServerModel true "被删除的服务地区信息（用服务实例信息包装）"
// @Success     200   {object} entity.ResultModel "返回true或false"
// @Router      /operator/CURD/deleteColony [delete]
func DeleteInstanceByColonyController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("DeleteInstanceByColonyController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	Assert, err1 := service.DeleteColonyServer(Server)
	if err1 != nil {
		util.Handle(err1, c)
		return
	}
	if Assert {
		c.JSON(http.StatusOK, entity.NewSuccessResult(Assert))
	} else {
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "命名空间或集群不存在"))
	}
}

// GetDeleteInstancesController
// @Summary     获取全部黑名单服务实例
// @Description 由管理者调用的获取当前全部黑名单服务实例
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200 {object} entity.ResultModel{data=[]entity.ServerModel} "返回黑名单中服务实例切片数组"
// @Router      /operator/CURD/getDeleteInstances [GET]
func GetDeleteInstancesController(c *gin.Context) {
	servers, err := service.GetBlacklistServer()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(servers))
}

// CancelDeleteInstancesController
// @Summary     删除黑名单中的实例信息
// @Description 由管理者调用删除黑名单中的实例信息。
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Param       Model body     entity.ServerModel true "从黑名单中清除的实例信息"
// @Success     200   {object} entity.ResultModel "返回true或false"
// @Router      /operator/CURD/cancelDeleteInstance [delete]
func CancelDeleteInstancesController(c *gin.Context) {
	Server := entity.NewServerModel()
	if err := c.BindJSON(Server); err != nil {
		exception.HandleException(exception.NewUserError("CancelDeleteInstancesController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	Assert1, err1 := service.CheckDeleteServer(Server)
	if err1 != nil {
		util.Handle(err1, c)
		return
	}
	if Assert1 {
		Assert2, err2 := service.DeleteInstanceFromBlacklist(Server)
		if err2 != nil {
			util.Handle(err2, c)
			return
		}
		c.JSON(http.StatusOK, entity.NewSuccessResult(Assert2))
	} else {
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "实例不在拒绝队列中"))
	}
}

// GetStatusController
// @Summary     获取服务器状态
// @Description 由管理员调用获取当前中心服务器状态
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200 {object} entity.ResultModel{data=entity.ComputerInfoModel} "返回电脑状态"
// @Router      /operator/cluster/getStatus [GET]
func GetStatusController(c *gin.Context) {
	computer, err := service.GetCenterStatus()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(computer))
}

// GetClusterLeaderController
// @Summary     获取中心集群领导者
// @Description 由管理员调用获取中心集群领导者
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200 {object} entity.ResultModel{} "返回中心集群领导者名称"
// @Router      /operator/cluster/getClusterLeader [GET]
func GetClusterLeaderController(c *gin.Context) {
	leader, err := service.GetClusterLeader()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(leader))
}

// GetClusterStatusController
// @Summary     获取中心当前身份状态
// @Description 由管理员调用获取中心当前身份状态
// @Tags        管理层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Success     200 {object} entity.ResultModel{data=syncBean.StatusLevel} "返回集群状态"
// @Router      /operator/cluster/getClusterStatus [GET]
func GetClusterStatusController(c *gin.Context) {
	leader, err := service.GetClusterStatus()
	if err != nil {
		util.Handle(err, c)
		return
	}
	c.JSON(http.StatusOK, entity.NewSuccessResult(leader))
}

// RootController
// @Summary     管理员登录
// @Description 管理员登录接口。
// @Tags        服务层
// @Accept      application/json
// @Produce     application/json
// @Security    ApiKeyAuth
// @Param       Model body     entity.Root true "管理员账号密码"
// @Success     200   {object} entity.ResultModel "返回token"
// @Router      /message/manager/login [POST]
func RootController(c *gin.Context) {
	root := entity.NewRootModel()
	err := c.BindJSON(&root)
	if err != nil {
		exception.HandleException(exception.NewUserError("GetServersController", "参数绑定错误-"+err.Error()))
		c.JSON(http.StatusOK, entity.NewFalseResult("false", "参数绑定错误-"+err.Error()))
		return
	}
	account := encryption.Base64Decode(root.Account)
	password := encryption.Base64Decode(root.Password)
	if (account == config.Root.RootAccount) && (password == config.Root.RootPassword) {
		t, err := token.GenerateToken(config.Root.RootAccount, config.Root.RootPassword)
		if err != nil {
			util.Handle(err, c)
			return
		}
		c.JSON(http.StatusOK, entity.NewSuccessResult(struct {
			Token string `json:"token"`
		}{Token: t}))
		return
	}
	c.JSON(http.StatusOK, entity.NewFalseResult("false", "密码错误"))
}
