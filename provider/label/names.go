package label

// Traefik labels
const (
	Prefix                                         = "traefik."
	SuffixBackend                                  = "backend"
	SuffixDomain                                   = "domain"
	SuffixEnable                                   = "enable"
	SuffixPort                                     = "port"
	SuffixPortIndex                                = "portIndex"
	SuffixProtocol                                 = "protocol"
	SuffixTags                                     = "tags"
	SuffixWeight                                   = "weight"
	SuffixBackendID                                = "backend.id"
	SuffixBackendCircuitBreaker                    = "backend.circuitbreaker"
	SuffixBackendCircuitBreakerExpression          = "backend.circuitbreaker.expression"
	SuffixBackendHealthCheckPath                   = "backend.healthcheck.path"
	SuffixBackendHealthCheckInterval               = "backend.healthcheck.interval"
	SuffixBackendLoadBalancerMethod                = "backend.loadbalancer.method"
	SuffixBackendLoadBalancerSticky                = "backend.loadbalancer.sticky"
	SuffixBackendLoadBalancerStickiness            = "backend.loadbalancer.stickiness"
	SuffixBackendLoadBalancerStickinessCookieName  = "backend.loadbalancer.stickiness.cookieName"
	SuffixBackendMaxConnAmount                     = "backend.maxconn.amount"
	SuffixBackendMaxConnExtractorFunc              = "backend.maxconn.extractorfunc"
	SuffixFrontendAuthBasic                        = "frontend.auth.basic"
	SuffixFrontendBackend                          = "frontend.backend"
	SuffixFrontendEntryPoints                      = "frontend.entryPoints"
	SuffixFrontendRequestHeaders                   = "frontend.headers.customRequestHeaders"
	SuffixFrontendResponseHeaders                  = "frontend.headers.customResponseHeaders"
	SuffixFrontendHeadersAllowedHosts              = "frontend.headers.allowedHosts"
	SuffixFrontendHeadersHostsProxyHeaders         = "frontend.headers.hostsProxyHeaders"
	SuffixFrontendHeadersSSLRedirect               = "frontend.headers.SSLRedirect"
	SuffixFrontendHeadersSSLTemporaryRedirect      = "frontend.headers.SSLTemporaryRedirect"
	SuffixFrontendHeadersSSLHost                   = "frontend.headers.SSLHost"
	SuffixFrontendHeadersSSLProxyHeaders           = "frontend.headers.SSLProxyHeaders"
	SuffixFrontendHeadersSTSSeconds                = "frontend.headers.STSSeconds"
	SuffixFrontendHeadersSTSIncludeSubdomains      = "frontend.headers.STSIncludeSubdomains"
	SuffixFrontendHeadersSTSPreload                = "frontend.headers.STSPreload"
	SuffixFrontendHeadersForceSTSHeader            = "frontend.headers.forceSTSHeader"
	SuffixFrontendHeadersFrameDeny                 = "frontend.headers.frameDeny"
	SuffixFrontendHeadersCustomFrameOptionsValue   = "frontend.headers.customFrameOptionsValue"
	SuffixFrontendHeadersContentTypeNosniff        = "frontend.headers.contentTypeNosniff"
	SuffixFrontendHeadersBrowserXSSFilter          = "frontend.headers.browserXSSFilter"
	SuffixFrontendHeadersContentSecurityPolicy     = "frontend.headers.contentSecurityPolicy"
	SuffixFrontendHeadersPublicKey                 = "frontend.headers.publicKey"
	SuffixFrontendHeadersReferrerPolicy            = "frontend.headers.referrerPolicy"
	SuffixFrontendHeadersIsDevelopment             = "frontend.headers.isDevelopment"
	SuffixFrontendPassHostHeader                   = "frontend.passHostHeader"
	SuffixFrontendPassTLSCert                      = "frontend.passTLSCert"
	SuffixFrontendPriority                         = "frontend.priority"
	SuffixFrontendRedirect                         = "frontend.redirect"
	SuffixFrontendRule                             = "frontend.rule"
	SuffixFrontendRuleType                         = "frontend.rule.type"
	SuffixFrontendWhitelistSourceRange             = "frontend.whitelistSourceRange"
	SuffixFrontendValue                            = "frontend.value"
	TraefikDomain                                  = Prefix + SuffixDomain
	TraefikEnable                                  = Prefix + SuffixEnable
	TraefikPort                                    = Prefix + SuffixPort
	TraefikPortIndex                               = Prefix + SuffixPortIndex
	TraefikProtocol                                = Prefix + SuffixProtocol
	TraefikTags                                    = Prefix + SuffixTags
	TraefikWeight                                  = Prefix + SuffixWeight
	TraefikBackend                                 = Prefix + SuffixBackend
	TraefikBackendID                               = Prefix + SuffixBackendID
	TraefikBackendCircuitBreaker                   = Prefix + SuffixBackendCircuitBreaker
	TraefikBackendCircuitBreakerExpression         = Prefix + SuffixBackendCircuitBreakerExpression
	TraefikBackendHealthCheckPath                  = Prefix + SuffixBackendHealthCheckPath
	TraefikBackendHealthCheckInterval              = Prefix + SuffixBackendHealthCheckInterval
	TraefikBackendLoadBalancerMethod               = Prefix + SuffixBackendLoadBalancerMethod
	TraefikBackendLoadBalancerSticky               = Prefix + SuffixBackendLoadBalancerSticky
	TraefikBackendLoadBalancerStickiness           = Prefix + SuffixBackendLoadBalancerStickiness
	TraefikBackendLoadBalancerStickinessCookieName = Prefix + SuffixBackendLoadBalancerStickinessCookieName
	TraefikBackendMaxConnAmount                    = Prefix + SuffixBackendMaxConnAmount
	TraefikBackendMaxConnExtractorFunc             = Prefix + SuffixBackendMaxConnExtractorFunc
	TraefikFrontendAuthBasic                       = Prefix + SuffixFrontendAuthBasic
	TraefikFrontendEntryPoints                     = Prefix + SuffixFrontendEntryPoints
	TraefikFrontendPassHostHeader                  = Prefix + SuffixFrontendPassHostHeader
	TraefikFrontendPassTLSCert                     = Prefix + SuffixFrontendPassTLSCert
	TraefikFrontendPriority                        = Prefix + SuffixFrontendPriority
	TraefikFrontendRule                            = Prefix + SuffixFrontendRule
	TraefikFrontendRuleType                        = Prefix + SuffixFrontendRuleType
	TraefikFrontendRedirect                        = Prefix + SuffixFrontendRedirect
	TraefikFrontendValue                           = Prefix + SuffixFrontendValue
	TraefikFrontendWhitelistSourceRange            = Prefix + SuffixFrontendWhitelistSourceRange
	TraefikFrontendRequestHeaders                  = Prefix + SuffixFrontendRequestHeaders
	TraefikFrontendResponseHeaders                 = Prefix + SuffixFrontendResponseHeaders
	TraefikFrontendAllowedHosts                    = Prefix + SuffixFrontendHeadersAllowedHosts
	TraefikFrontendHostsProxyHeaders               = Prefix + SuffixFrontendHeadersHostsProxyHeaders
	TraefikFrontendSSLRedirect                     = Prefix + SuffixFrontendHeadersSSLRedirect
	TraefikFrontendSSLTemporaryRedirect            = Prefix + SuffixFrontendHeadersSSLTemporaryRedirect
	TraefikFrontendSSLHost                         = Prefix + SuffixFrontendHeadersSSLHost
	TraefikFrontendSSLProxyHeaders                 = Prefix + SuffixFrontendHeadersSSLProxyHeaders
	TraefikFrontendSTSSeconds                      = Prefix + SuffixFrontendHeadersSTSSeconds
	TraefikFrontendSTSIncludeSubdomains            = Prefix + SuffixFrontendHeadersSTSIncludeSubdomains
	TraefikFrontendSTSPreload                      = Prefix + SuffixFrontendHeadersSTSPreload
	TraefikFrontendForceSTSHeader                  = Prefix + SuffixFrontendHeadersForceSTSHeader
	TraefikFrontendFrameDeny                       = Prefix + SuffixFrontendHeadersFrameDeny
	TraefikFrontendCustomFrameOptionsValue         = Prefix + SuffixFrontendHeadersCustomFrameOptionsValue
	TraefikFrontendContentTypeNosniff              = Prefix + SuffixFrontendHeadersContentTypeNosniff
	TraefikFrontendBrowserXSSFilter                = Prefix + SuffixFrontendHeadersBrowserXSSFilter
	TraefikFrontendContentSecurityPolicy           = Prefix + SuffixFrontendHeadersContentSecurityPolicy
	TraefikFrontendPublicKey                       = Prefix + SuffixFrontendHeadersPublicKey
	TraefikFrontendReferrerPolicy                  = Prefix + SuffixFrontendHeadersReferrerPolicy
	TraefikFrontendIsDevelopment                   = Prefix + SuffixFrontendHeadersIsDevelopment
)
