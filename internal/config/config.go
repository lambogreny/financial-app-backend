package config

import "github.com/namsral/flag"

var DataDirectory = flag.String("data-directory", "", "Path for loading template and migratuib scripts")
