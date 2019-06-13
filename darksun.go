package darksun

// Module is a representation of some part of the system that
// can be configured dark or light.
type Module interface {
	Dark() error
	Light() error
}
