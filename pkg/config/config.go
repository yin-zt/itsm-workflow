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
	UrlPathPrefix  = "itsm-workflow"
	CONST_SYS_DIRS = []string{StorageDir}
)

const (
	ServicePort = 8888
	StorageDir  = "/var/loger"
)
