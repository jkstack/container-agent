package core

type Agent struct {
	cfgDir  string
	cfg     *Configure
	version string
}

func New(dir, version string) *Agent {
	return &Agent{
		cfgDir:  dir,
		cfg:     load(dir),
		version: version,
	}
}

func (agent *Agent) AgentName() string {
	return "container-agent"
}

func (agent *Agent) Version() string {
	return agent.version
}
