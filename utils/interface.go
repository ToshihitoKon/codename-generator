package utils

type CodenameGenerator interface {
	GenerateCodename(string) (string, error)
}
