package walmart_client

const (
	RequestAuthContentType = "application/x-www-form-urlencoded"
	RequestContentType     = "application/json"
	RequestAccept          = "application/json"

	BaseUrlProd    = "https://marketplace.walmartapis.com/"
	BaseUrlSandbox = "https://sandbox.walmartapis.com"

	GetAccessTokenUri = "v3/token"
)

type Config struct {
	SandboxMode   bool
	WMServiceName string
	ClientID      string
	ClientSecret  string
}
