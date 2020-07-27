package runner

type ICmdRunner interface {
	Run() error
	String() string
}
