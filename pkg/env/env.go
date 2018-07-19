package env

type Env struct {
	IO *IO
}

func New(ioWantedLevel int) *Env {
	return &Env{
		IO: newIO(ioWantedLevel),
	}
}