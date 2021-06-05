package task

type Configuration struct {
	Config Config `yaml:"config"`
}

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type Step struct {
	Name     string    `yaml:"name,omitempty"`
	Commands []Command `yaml:"commands,omitempty"`
	Steps    []Step    `yaml:"steps,omitempty"`
}

type Config struct {
	Step []Step `yaml:"steps"`
}
