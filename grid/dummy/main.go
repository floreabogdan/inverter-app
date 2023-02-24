package dummy

type Dummy struct {
	config string
}

func New(config string) *Dummy {
	return &Dummy{
		config: config,
	}
}

func (d Dummy) SetConfig(config string) {
	print(config)
}

func (d Dummy) GetConfig() string {
	return "Hello world!"
}
