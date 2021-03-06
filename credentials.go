package sftp_server

type AuthenticationRequest struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type AuthenticationResponse struct {
	Server      string   `json:"server"`
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
}

type InvalidCredentialsError struct {
}

func (ice InvalidCredentialsError) Error() string {
	return "the credentials provided were invalid"
}

func IsInvalidCredentialsError(err error) bool {
	_, ok := err.(*InvalidCredentialsError)

	return ok
}