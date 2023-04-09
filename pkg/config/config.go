package config

var (
	LogOperateConfigStr = `
<seelog type="asynctimer" asyncinterval="1000" minlevel="trace" maxlevel="error">  
	<outputs formatid="common">  
		<buffered formatid="common" size="1048576" flushperiod="1000">  
			<rollingfile type="size" filename="/var/loger/itsm-workflow-operate.log" maxsize="104857600" maxrolls="5"/>  
		</buffered>
	</outputs>  	  
	 <formats>
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />  
	 </formats>  
</seelog>
`
	LogAccessConfigStr = `
<seelog type="asynctimer" asyncinterval="1000" minlevel="trace" maxlevel="error">  
	<outputs formatid="common">  
		<buffered formatid="common" size="1048576" flushperiod="1000">  
			<rollingfile type="size" filename="/var/loger/itsm-workflow-access.log" maxsize="104857600" maxrolls="3"/>  
		</buffered>
	</outputs> 
	 <formats>
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />  
	 </formats>  
</seelog>
`
	UrlPathPrefix  = "itsm-workflow"      // 接口调用前缀
	ConstSysDirs   = []string{StorageDir} // 需要创建的目录
	MysqlUsername  = "root"               // 连接mysql的用户
	MysqlPassword  = ""
	MysqlHost      = ""
	MysqlPort      = 306
	MysqlDatabase  = "itsm-workflow"
	MysqlCharset   = "utf8mb4"                                  // 编码方式
	MysqlCollation = "utf8mb4_general_ci"                       // 字符集(utf8mb4_general_ci速度比utf8mb4_unicode_ci快些)
	MysqlQuery     = "parseTime=True&loc=Local&timeout=10000ms" //连接字符串参数
	MysqlLogMode   = true                                       // 是否打印日志
)

const (
	ServicePort = 8888
	StorageDir  = "/var/loger"
)
