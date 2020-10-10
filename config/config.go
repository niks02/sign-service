package config

type configData struct {
	Port   int          `json:"port"`
	Crypto cryptoConfig `json:"crypto"`
}

type cryptoConfig struct {
	PublicKey  string `json:"public"`
	PrivateKey string `json:"private"`
	Algo       string `json:"algo"`
}

// GetPort returns port
func (c *configData) GetPort() int {
	return c.Port
}

// GetPublicKey returns publicKey
func (c *configData) GetPublicKey() string {
	return c.Crypto.PublicKey
}

// GetPrivateKey returns privateKey
func (c *configData) GetPrivateKey() string {
	return c.Crypto.PrivateKey
}

// GetSigningAlgo returns signingAlgo
func (c *configData) GetSigningAlgo() string {
	return c.Crypto.Algo
}
