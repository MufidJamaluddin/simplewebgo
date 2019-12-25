package model

// CredentialModel tipe data saat user login
type CredentialModel struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Group    string `json:"Group"`
}
