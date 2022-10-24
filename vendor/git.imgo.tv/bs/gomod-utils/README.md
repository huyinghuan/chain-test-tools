

# goutils
`import "git.imgo.tv/bs/go-sdk/goutils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func CreateHTTPClient(timeout time.Duration) *http.Client](#CreateHTTPClient)
* [func InitLog(log LogInterface) error](#InitLog)
* [func InitSignal(signalReload SignalReload)](#InitSignal)
* [type HTTPConnectionPool](#HTTPConnectionPool)
  * [func NewHTTPConnectionPool(timeout time.Duration, poolNum int) *HTTPConnectionPool](#NewHTTPConnectionPool)
  * [func (cp *HTTPConnectionPool) BatchRequest(httpDatas []*HTTPData)](#HTTPConnectionPool.BatchRequest)
  * [func (cp *HTTPConnectionPool) Request(request *http.Request) (*http.Response, error)](#HTTPConnectionPool.Request)
  * [func (cp *HTTPConnectionPool) SetName(name string)](#HTTPConnectionPool.SetName)
  * [func (cp *HTTPConnectionPool) Status() string](#HTTPConnectionPool.Status)
* [type HTTPData](#HTTPData)
  * [func NewHTTPData(request *http.Request) *HTTPData](#NewHTTPData)
  * [func NewHTTPDataWithExtra(request *http.Request, extra interface{}) *HTTPData](#NewHTTPDataWithExtra)
* [type LogInterface](#LogInterface)
* [type Redis](#Redis)
  * [func NewRedis(redisConfig RedisConfig) *Redis](#NewRedis)
  * [func (r *Redis) Client() redis.Cmdable](#Redis.Client)
  * [func (r *Redis) Del(key ...string) error](#Redis.Del)
  * [func (r *Redis) Exists(key string) (bool, error)](#Redis.Exists)
  * [func (r *Redis) Get(key string) (string, error)](#Redis.Get)
  * [func (r *Redis) HDel(key, field string) error](#Redis.HDel)
  * [func (r *Redis) HGet(key, field string) (string, error)](#Redis.HGet)
  * [func (r *Redis) HGetAllMap(key string) (map[string]string, error)](#Redis.HGetAllMap)
  * [func (r *Redis) HMSet(key string, fields map[string]string) error](#Redis.HMSet)
  * [func (r *Redis) HSet(key, feild, value string) error](#Redis.HSet)
  * [func (r *Redis) Incby(key string, value int64) (int64, error)](#Redis.Incby)
  * [func (r *Redis) Keys(pattern string) ([]string, error)](#Redis.Keys)
  * [func (r *Redis) Mget(keys ...string) ([]interface{}, error)](#Redis.Mget)
  * [func (r *Redis) PSubscribe(channels ...string) (*redis.PubSub, error)](#Redis.PSubscribe)
  * [func (r *Redis) Publish(channel, message string) error](#Redis.Publish)
  * [func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error](#Redis.Set)
  * [func (r *Redis) SetExpire(key string, expiration time.Duration) error](#Redis.SetExpire)
  * [func (r *Redis) SetNx(key string, value interface{}, expiration time.Duration) (bool, error)](#Redis.SetNx)
  * [func (r *Redis) Subscribe(channel string) (*redis.PubSub, error)](#Redis.Subscribe)
  * [func (r *Redis) TTL(key string) (time.Duration, error)](#Redis.TTL)
* [type RedisConfig](#RedisConfig)
* [type ServerContext](#ServerContext)
  * [func NewContext(msg string) *ServerContext](#NewContext)
  * [func (sc *ServerContext) AddNotes(key string, val interface{})](#ServerContext.AddNotes)
  * [func (sc *ServerContext) Critical(format string, args ...interface{})](#ServerContext.Critical)
  * [func (sc *ServerContext) Debug(format string, args ...interface{})](#ServerContext.Debug)
  * [func (sc *ServerContext) Error(format string, args ...interface{})](#ServerContext.Error)
  * [func (sc *ServerContext) Flush()](#ServerContext.Flush)
  * [func (sc *ServerContext) GetUUID() string](#ServerContext.GetUUID)
  * [func (sc *ServerContext) Info(format string, args ...interface{})](#ServerContext.Info)
  * [func (sc *ServerContext) Notice(format string, args ...interface{})](#ServerContext.Notice)
  * [func (sc *ServerContext) SetUUID(uuid string)](#ServerContext.SetUUID)
  * [func (sc *ServerContext) StartTimer()](#ServerContext.StartTimer)
  * [func (sc *ServerContext) StopTimer(key string)](#ServerContext.StopTimer)
  * [func (sc *ServerContext) Warning(format string, args ...interface{})](#ServerContext.Warning)
* [type SignalReload](#SignalReload)


#### <a name="pkg-files">Package files</a>
[context.go](/src/git.imgo.tv/bs/go-sdk/goutils/context.go) [http.go](/src/git.imgo.tv/bs/go-sdk/goutils/http.go) [log.go](/src/git.imgo.tv/bs/go-sdk/goutils/log.go) [redis.go](/src/git.imgo.tv/bs/go-sdk/goutils/redis.go) [signals.go](/src/git.imgo.tv/bs/go-sdk/goutils/signals.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    //Log 日志
    Log = logging.MustGetLogger("mgtv")
)
```


## <a name="CreateHTTPClient">func</a> [CreateHTTPClient](/http.go?s=362:419#L19)
``` go
func CreateHTTPClient(timeout time.Duration) *http.Client
```
CreateHTTPClient 创建 httpClient



## <a name="InitLog">func</a> [InitLog](/log.go?s=3468:3504#L125)
``` go
func InitLog(log LogInterface) error
```
InitLog 日志初始化



## <a name="InitSignal">func</a> [InitSignal](/signals.go?s=345:387#L20)
``` go
func InitSignal(signalReload SignalReload)
```
InitSignal 用户信号量初始化

SIGUSR1: 日志文件重新打开类似 nginx -s reload, 完成日志切割

SIGUSR2: 配置文件重新加载,可以完成比如日志级别的动态改变




## <a name="HTTPConnectionPool">type</a> [HTTPConnectionPool](/http.go?s=1269:1595#L50)
``` go
type HTTPConnectionPool struct {
    // contains filtered or unexported fields
}
```
HTTPConnectionPool http连接池







### <a name="NewHTTPConnectionPool">func</a> [NewHTTPConnectionPool](/http.go?s=1647:1729#L62)
``` go
func NewHTTPConnectionPool(timeout time.Duration, poolNum int) *HTTPConnectionPool
```
NewHTTPConnectionPool http连接池构造函数





### <a name="HTTPConnectionPool.BatchRequest">func</a> (\*HTTPConnectionPool) [BatchRequest](/http.go?s=3170:3235#L124)
``` go
func (cp *HTTPConnectionPool) BatchRequest(httpDatas []*HTTPData)
```
BatchRequest http批量请求接口




### <a name="HTTPConnectionPool.Request">func</a> (\*HTTPConnectionPool) [Request](/http.go?s=2658:2742#L105)
``` go
func (cp *HTTPConnectionPool) Request(request *http.Request) (*http.Response, error)
```
Request http请求接口




### <a name="HTTPConnectionPool.SetName">func</a> (\*HTTPConnectionPool) [SetName](/http.go?s=2276:2326#L87)
``` go
func (cp *HTTPConnectionPool) SetName(name string)
```
SetName 设置连接池名字，方便统计




### <a name="HTTPConnectionPool.Status">func</a> (\*HTTPConnectionPool) [Status](/http.go?s=3779:3824#L148)
``` go
func (cp *HTTPConnectionPool) Status() string
```
Status 获取连接池状态并初始化状态




## <a name="HTTPData">type</a> [HTTPData](/http.go?s=637:805#L31)
``` go
type HTTPData struct {
    Request   *http.Request
    Response  *http.Response
    Err       error
    ExtraData interface{} //http 请求的自定义信息
    // contains filtered or unexported fields
}
```
HTTPData http请求和响应







### <a name="NewHTTPData">func</a> [NewHTTPData](/http.go?s=842:891#L40)
``` go
func NewHTTPData(request *http.Request) *HTTPData
```
NewHTTPData HTTPData constructor


### <a name="NewHTTPDataWithExtra">func</a> [NewHTTPDataWithExtra](/http.go?s=1045:1122#L45)
``` go
func NewHTTPDataWithExtra(request *http.Request, extra interface{}) *HTTPData
```
NewHTTPDataWithExtra HTTPData constructor with extra data





## <a name="LogInterface">type</a> [LogInterface](/log.go?s=160:234#L12)
``` go
type LogInterface interface {
    GetLogPath() string
    GetLogLevel() string
}
```
LogInterface log日志接口，实现get path和get log level方法










## <a name="Redis">type</a> [Redis](/redis.go?s=299:342#L21)
``` go
type Redis struct {
    // contains filtered or unexported fields
}
```






### <a name="NewRedis">func</a> [NewRedis](/redis.go?s=344:389#L25)
``` go
func NewRedis(redisConfig RedisConfig) *Redis
```




### <a name="Redis.Client">func</a> (\*Redis) [Client](/redis.go?s=7104:7142#L269)
``` go
func (r *Redis) Client() redis.Cmdable
```



### <a name="Redis.Del">func</a> (\*Redis) [Del](/redis.go?s=5733:5773#L220)
``` go
func (r *Redis) Del(key ...string) error
```



### <a name="Redis.Exists">func</a> (\*Redis) [Exists](/redis.go?s=5011:5059#L195)
``` go
func (r *Redis) Exists(key string) (bool, error)
```



### <a name="Redis.Get">func</a> (\*Redis) [Get](/redis.go?s=1961:2008#L78)
``` go
func (r *Redis) Get(key string) (string, error)
```
get string key just for freq condition




### <a name="Redis.HDel">func</a> (\*Redis) [HDel](/redis.go?s=3896:3941#L150)
``` go
func (r *Redis) HDel(key, field string) error
```



### <a name="Redis.HGet">func</a> (\*Redis) [HGet](/redis.go?s=3440:3495#L132)
``` go
func (r *Redis) HGet(key, field string) (string, error)
```



### <a name="Redis.HGetAllMap">func</a> (\*Redis) [HGetAllMap](/redis.go?s=4106:4171#L158)
``` go
func (r *Redis) HGetAllMap(key string) (map[string]string, error)
```



### <a name="Redis.HMSet">func</a> (\*Redis) [HMSet](/redis.go?s=3210:3275#L123)
``` go
func (r *Redis) HMSet(key string, fields map[string]string) error
```



### <a name="Redis.HSet">func</a> (\*Redis) [HSet](/redis.go?s=4566:4618#L177)
``` go
func (r *Redis) HSet(key, feild, value string) error
```



### <a name="Redis.Incby">func</a> (\*Redis) [Incby](/redis.go?s=5269:5330#L205)
``` go
func (r *Redis) Incby(key string, value int64) (int64, error)
```
IncrBy(key string, value int64) *IntCmd




### <a name="Redis.Keys">func</a> (\*Redis) [Keys](/redis.go?s=5933:5987#L228)
``` go
func (r *Redis) Keys(pattern string) ([]string, error)
```



### <a name="Redis.Mget">func</a> (\*Redis) [Mget](/redis.go?s=2963:3022#L114)
``` go
func (r *Redis) Mget(keys ...string) ([]interface{}, error)
```



### <a name="Redis.PSubscribe">func</a> (\*Redis) [PSubscribe](/redis.go?s=6694:6763#L255)
``` go
func (r *Redis) PSubscribe(channels ...string) (*redis.PubSub, error)
```



### <a name="Redis.Publish">func</a> (\*Redis) [Publish](/redis.go?s=6397:6451#L245)
``` go
func (r *Redis) Publish(channel, message string) error
```
Publish 由于publish 不是基本命令，需要特殊实现




### <a name="Redis.Set">func</a> (\*Redis) [Set](/redis.go?s=2401:2483#L96)
``` go
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error
```



### <a name="Redis.SetExpire">func</a> (\*Redis) [SetExpire](/redis.go?s=4780:4849#L186)
``` go
func (r *Redis) SetExpire(key string, expiration time.Duration) error
```



### <a name="Redis.SetNx">func</a> (\*Redis) [SetNx](/redis.go?s=2671:2763#L105)
``` go
func (r *Redis) SetNx(key string, value interface{}, expiration time.Duration) (bool, error)
```



### <a name="Redis.Subscribe">func</a> (\*Redis) [Subscribe](/redis.go?s=6904:6968#L262)
``` go
func (r *Redis) Subscribe(channel string) (*redis.PubSub, error)
```



### <a name="Redis.TTL">func</a> (\*Redis) [TTL](/redis.go?s=5503:5557#L213)
``` go
func (r *Redis) TTL(key string) (time.Duration, error)
```



## <a name="RedisConfig">type</a> [RedisConfig](/redis.go?s=80:297#L11)
``` go
type RedisConfig interface {
    GetPassword() string
    GetAddr() string
    GetPoolNum() int
    GetReadTimeout() time.Duration
    GetWriteTimeout() time.Duration
    GetPoolTimeout() time.Duration
    GetDialTimeout() time.Duration
}
```









## <a name="ServerContext">type</a> [ServerContext](/context.go?s=144:288#L15)
``` go
type ServerContext struct {
    // contains filtered or unexported fields
}
```
ServerContext 日志上下文







### <a name="NewContext">func</a> [NewContext](/context.go?s=317:359#L24)
``` go
func NewContext(msg string) *ServerContext
```
NewContext 构造函数





### <a name="ServerContext.AddNotes">func</a> (\*ServerContext) [AddNotes](/context.go?s=1242:1304#L62)
``` go
func (sc *ServerContext) AddNotes(key string, val interface{})
```
AddNotes 添加kv对到日志中




### <a name="ServerContext.Critical">func</a> (\*ServerContext) [Critical](/context.go?s=2725:2794#L110)
``` go
func (sc *ServerContext) Critical(format string, args ...interface{})
```
Critical Critical日志




### <a name="ServerContext.Debug">func</a> (\*ServerContext) [Debug](/context.go?s=1699:1765#L78)
``` go
func (sc *ServerContext) Debug(format string, args ...interface{})
```
Debug debug日志




### <a name="ServerContext.Error">func</a> (\*ServerContext) [Error](/context.go?s=2471:2537#L103)
``` go
func (sc *ServerContext) Error(format string, args ...interface{})
```
Error Error日志




### <a name="ServerContext.Flush">func</a> (\*ServerContext) [Flush](/context.go?s=1462:1494#L69)
``` go
func (sc *ServerContext) Flush()
```
Flush flush所有AddNotes日志，通常工作流结束调用




### <a name="ServerContext.GetUUID">func</a> (\*ServerContext) [GetUUID](/context.go?s=754:795#L44)
``` go
func (sc *ServerContext) GetUUID() string
```
GetUUID 获取当前上下文uuid




### <a name="ServerContext.Info">func</a> (\*ServerContext) [Info](/context.go?s=1869:1934#L84)
``` go
func (sc *ServerContext) Info(format string, args ...interface{})
```
Info Info日志




### <a name="ServerContext.Notice">func</a> (\*ServerContext) [Notice](/context.go?s=2041:2108#L90)
``` go
func (sc *ServerContext) Notice(format string, args ...interface{})
```
Notice Notice日志




### <a name="ServerContext.SetUUID">func</a> (\*ServerContext) [SetUUID](/context.go?s=591:636#L35)
``` go
func (sc *ServerContext) SetUUID(uuid string)
```
SetUUID 设置上下文uuid，用于trace整个工作流




### <a name="ServerContext.StartTimer">func</a> (\*ServerContext) [StartTimer](/context.go?s=904:941#L49)
``` go
func (sc *ServerContext) StartTimer()
```
StartTimer 调用开始计时，用于统计程序耗时，和StopTimer配合使用




### <a name="ServerContext.StopTimer">func</a> (\*ServerContext) [StopTimer](/context.go?s=1024:1070#L54)
``` go
func (sc *ServerContext) StopTimer(key string)
```
StopTimer 结束计时，和StartTimer配合使用




### <a name="ServerContext.Warning">func</a> (\*ServerContext) [Warning](/context.go?s=2219:2287#L96)
``` go
func (sc *ServerContext) Warning(format string, args ...interface{})
```
Warning Warning日志




## <a name="SignalReload">type</a> [SignalReload](/signals.go?s=96:137#L11)
``` go
type SignalReload interface {
    Reload()
}
```
SignalReload interface














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
