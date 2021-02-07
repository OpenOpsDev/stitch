package configs

type DependencyConfig struct {
	PackageManager string `yaml:"package_manager"`
	Dependencies []string `yaml:"dependencies"`
}

func (d *DependencyConfig) Install() {
	//  exec.Cmd(d.PackageManager)
}