package main

type Config struct {
	OrganizeByType bool
	OrganizeByDate bool
	OrganizeBySize bool
	OrganizeByName bool
	SkipHidden bool
	TargetPath string
}

func DefaultConfig() Config{
	return Config{
		OrganizeByType: true,
		OrganizeByDate: false,
		OrganizeBySize: false,
		OrganizeByName: false,
		SkipHidden: true,
		TargetPath: ".",
	}
}

