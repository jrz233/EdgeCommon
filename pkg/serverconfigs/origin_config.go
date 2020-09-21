package serverconfigs

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/sslconfigs"
	"net"
	"strconv"
	"strings"
	"time"
)

// 源站服务配置
type OriginConfig struct {
	Id          int64                 `yaml:"id" json:"id"`                   // ID
	IsOn        bool                  `yaml:"isOn" json:"isOn"`               // 是否启用 TODO
	Version     int                   `yaml:"version" json:"version"`         // 版本
	Name        string                `yaml:"name" json:"name"`               // 名称 TODO
	Addr        *NetworkAddressConfig `yaml:"addr" json:"addr"`               // 地址
	Description string                `yaml:"description" json:"description"` // 描述 TODO
	Code        string                `yaml:"code" json:"code"`               // 代号 TODO

	Weight       uint                 `yaml:"weight" json:"weight"`           // 权重 TODO
	ConnTimeout  *shared.TimeDuration `yaml:"failTimeout" json:"failTimeout"` // 连接失败超时 TODO
	ReadTimeout  *shared.TimeDuration `yaml:"readTimeout" json:"readTimeout"` // 读取超时时间 TODO
	IdleTimeout  *shared.TimeDuration `yaml:"idleTimeout" json:"idleTimeout"` // 空闲连接超时时间 TODO
	MaxFails     int                  `yaml:"maxFails" json:"maxFails"`       // 最多失败次数 TODO
	MaxConns     int                  `yaml:"maxConns" json:"maxConns"`       // 最大并发连接数 TODO
	MaxIdleConns int                  `yaml:"idleConns" json:"idleConns"`     // 最大空闲连接数 TODO

	RequestURI string `yaml:"requestURI" json:"requestURI"` // 转发后的请求URI TODO
	Host       string `yaml:"host" json:"host"`             // 自定义主机名 TODO

	RequestHeaders  *shared.HTTPHeaderPolicy `yaml:"requestHeaders" json:"requestHeaders"` // 请求Header设置 TODO
	ResponseHeaders *shared.HTTPHeaderPolicy `yaml:"responseHeaders" json:"responseHeaders"`

	// 健康检查URL，目前支持：
	// - http|https 返回2xx-3xx认为成功
	HealthCheck *HealthCheckConfig `yaml:"healthCheck" json:"healthCheck"`

	Cert *sslconfigs.SSLCertConfig `yaml:"cert" json:"cert"` // 请求源服务器用的证书

	// ftp
	FTP *OriginFTPConfig `yaml:"ftp" json:"ftp"`

	connTimeoutDuration time.Duration
	readTimeoutDuration time.Duration
	idleTimeoutDuration time.Duration

	hasRequestURI bool
	requestPath   string
	requestArgs   string

	hasRequestHeaders  bool
	hasResponseHeaders bool

	hasHost bool

	uniqueKey string

	hasAddrVariables bool // 地址中是否含有变量
}

// 校验
func (this *OriginConfig) Init() error {
	// 证书
	if this.Cert != nil {
		err := this.Cert.Init()
		if err != nil {
			return err
		}
	}

	// unique key
	this.uniqueKey = strconv.FormatInt(this.Id, 10) + "@" + fmt.Sprintf("%d", this.Version)

	// failTimeout
	this.connTimeoutDuration = this.ConnTimeout.Duration()

	// readTimeout
	this.readTimeoutDuration = this.ReadTimeout.Duration()

	// idleTimeout
	this.idleTimeoutDuration = this.IdleTimeout.Duration()

	// Headers
	if this.RequestHeaders != nil {
		err := this.RequestHeaders.Init()
		if err != nil {
			return err
		}
	}
	if this.ResponseHeaders != nil {
		err := this.ResponseHeaders.Init()
		if err != nil {
			return err
		}
	}

	// request uri
	if len(this.RequestURI) == 0 || this.RequestURI == "${requestURI}" {
		this.hasRequestURI = false
	} else {
		this.hasRequestURI = true

		if strings.Contains(this.RequestURI, "?") {
			pieces := strings.SplitN(this.RequestURI, "?", -1)
			this.requestPath = pieces[0]
			this.requestArgs = pieces[1]
		} else {
			this.requestPath = this.RequestURI
		}
	}

	// TODO init health check

	// headers
	this.hasRequestHeaders = this.RequestHeaders != nil && !this.RequestHeaders.IsEmpty()
	this.hasRequestHeaders = this.ResponseHeaders != nil && !this.ResponseHeaders.IsEmpty()

	// host
	this.hasHost = len(this.Host) > 0

	// variables
	// TODO 在host和port中支持变量
	this.hasAddrVariables = false

	return nil
}

// 候选对象代号
func (this *OriginConfig) CandidateCodes() []string {
	codes := []string{strconv.FormatInt(this.Id, 10)}
	if len(this.Code) > 0 {
		codes = append(codes, this.Code)
	}
	return codes
}

// 候选对象权重
func (this *OriginConfig) CandidateWeight() uint {
	return this.Weight
}

// 连接源站
func (this *OriginConfig) Connect() (net.Conn, error) {
	if this.Addr == nil {
		return nil, errors.New("origin server address should not be empty")
	}

	switch this.Addr.Protocol {
	case "", ProtocolTCP:
		// TODO 支持TCP4/TCP6
		// TODO 支持指定特定网卡
		// TODO Addr支持端口范围，如果有多个端口时，随机一个端口使用
		return net.DialTimeout("tcp", this.Addr.Host+":"+this.Addr.PortRange, this.connTimeoutDuration)
	case ProtocolTLS:
		// TODO 支持TCP4/TCP6
		// TODO 支持指定特定网卡
		// TODO Addr支持端口范围，如果有多个端口时，随机一个端口使用
		// TODO 支持使用证书
		return tls.Dial("tcp", this.Addr.Host+":"+this.Addr.PortRange, &tls.Config{})
	}

	// TODO 支持从Unix、Pipe、HTTP、HTTPS中读取数据

	return nil, errors.New("invalid scheme '" + this.Addr.Protocol.String() + "'")
}