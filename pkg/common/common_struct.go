package common


type PasswordEntry struct {
	ApplicationName     string `json:"applicationname, omitempty"`
	Username 			string `json:"username"`
	Password 			string `json:"password"`
}

type RootUserEntry struct{
	Username	string `json:"username"`
	Password	string `json:"password"`
}