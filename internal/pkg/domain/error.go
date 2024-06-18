package domain

const (
	criticalError = 2
	domainError   = 1
)

type BaseError interface {
	Error() string
	ExtraItems() map[string]interface{}
	Severity() int
}

type CriticalError struct {
}

func (d *CriticalError) Severity() int {
	return criticalError
}

type DomainError struct {
}

func (d *DomainError) Severity() int {
	return domainError
}
