package main

type CVGenCmd struct {
	generator CVGenerator
}

func NewCVGenCmd(conf string) (*CVGenCmd, error) {
	config, err := ParseConfigFile(conf)
	if err != nil {
		return nil, err
	}

	var gen CVGenerator
	switch config.Style.Mode {
	case ModeText:
		gen, err = NewTextCVGenerator(config)
		if err != nil {
			return nil, err
		}
	}
	return &CVGenCmd{
		generator: gen,
	}, nil
}

func (c *CVGenCmd) Run() error {
	return c.generator.Generate()
}
