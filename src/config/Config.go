package config

var (
	// Goroutine goroutine参数
	Goroutine struct {
		// MaxRoutineNum goroutine池最大线程数
		MaxRoutineNum int
		// CoreRoutineNum goroutine池核心线程数
		CoreRoutineNum int
		// RoutineTimeOut goroutine池线程超时时间
		RoutineTimeOut int
	}

	// Root root用户信息
	Root struct {
		// RootAccount 超级管理员账号
		RootAccount string
		// RootPassword 超级管理员密码
		RootPassword string
		// TokenEnable 是否开启token验证
		TokenEnable bool
		// TokenExpireTime 超级管理员token过期时间
		TokenExpireTime int
		// TokenSignKey 超级管理员token签名密钥
		TokenSignKey string
	}

	// Port 端口参数
	Port struct {
		// CenterPort 注册中心http端口
		CenterPort string
		// UDPTimeOut UDP超时时间
		UDPTimeOut int
	}

	// ServerRegister 服务注册参数
	ServerRegister struct {
		// ServerModelQueueNum 服务注册处理队列容量
		ServerModelQueueNum int

		// ServerModelHandleNum 服务注册处理器数量
		ServerModelHandleNum int
	}

	// ServerBeat 服务心跳参数
	ServerBeat struct {
		// ServerModelBeatEnable 服务注册心跳开关
		ServerModelBeatEnable bool

		// ServerModelBeatQueue 服务心跳处理队列容量
		ServerModelBeatQueue int

		// ServerBeatTime 服务心跳超时时间   单位：s
		ServerBeatTime int64
	}

	// CreateLeaderAlgorithm 记账人选举算法
	CreateLeaderAlgorithm string

	// ElectionTimeOut 选举超时时间
	ElectionTimeOut int64

	// Database 数据库参数
	Database struct {
		// DatabaseType 数据库类型
		DatabaseType string

		// DatabaseHost 数据库地址
		DatabaseHost string

		// DatabasePort 数据库端口
		DatabasePort string

		// DatabaseName 数据库名称
		DatabaseName string

		// DatabaseUser 数据库用户名
		DatabaseUser string

		// DatabasePassword 数据库密码
		DatabasePassword string

		// DatabaseMaxOpenConns 数据库最大连接数
		DatabaseMaxOpenConns int

		// DatabaseMaxIdleConns 数据库最大空闲连接数
		DatabaseMaxIdleConns int

		// DatabaseMaxLifetimeConns 数据库最大连接生命周期
		DatabaseMaxLifetimeConns int
	}

	// Persistence 持久化参数
	Persistence struct {
		// PersistenceEnable 是否开启持久化
		PersistenceEnable bool

		// PersistenceTime 持久化时间
		PersistenceTime int64

		// SoftDeleteEnable 是否开启软删除
		SoftDeleteEnable bool
	}

	// ListenTime 监听时间
	ListenTime int64

	// Cluster 集群参数
	Cluster struct {
		// ClusterEnable 是否开启集群
		ClusterEnable bool

		//TrackEnable 是否开启跟踪
		TrackEnable bool

		// ClusterName 集群名称
		ClusterName string

		//IP	地址
		IP string

		//Port	端口
		Port string

		//MaxTimeOut	最大超时时间
		MaxFollowTimeOut int64

		//MinTimeOut	最小超时时间
		MinFollowTimeOut int64

		//CandidateTimeOut	 最大候选人超时时间
		MaxCandidateTimeOut int64

		//MinCandidateTimeOut	最小候选人超时时间
		MinCandidateTimeOut int64

		//UDPTimeOut	UDP超时时间
		UDPTimeOut int64

		//UDPQueueNum	UDP队列
		UDPQueueNum int

		//LeaderSnapshotSyncTime	snapshot同步时间
		LeaderSnapshotSyncTime int64

		//LeaderHeartbeatTime		心跳间隔
		LeaderHeartbeatTime int64

		//LeaderQueueNum	leader队列
		LeaderQueueNum int

		//Clusters	集群
		Clusters []map[string]string

		//EnableEncryption	是否开启加密
		EnableEncryption bool

		//EncryptionKey	加密密钥
		EncryptionKey []byte
	}
)
