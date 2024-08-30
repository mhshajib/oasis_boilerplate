package sms

type SMS interface {
	Send(to, message string) (string, error)
}
