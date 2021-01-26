package defaultinterface

//PasswordService default interface
type PasswordService interface {
	Generate(raw string) (string, error)
	Compare(p1, p2 string) error
}
