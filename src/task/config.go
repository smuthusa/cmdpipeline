package task

type Configuration struct {
	Config Config `yaml:"config"`
}

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type Step struct {
	Name     string    `yaml:"name"`
	Commands []Command `yaml:"commands"`
}

type Stage struct {
	Name  string   `yaml:"name"`
	Steps []string `yaml:"steps"`
}

type Config struct {
	Step  []Step  `yaml:"steps"`
	Stage []Stage `yaml:"stages"`
}
