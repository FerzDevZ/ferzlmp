package vhost

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddVHost(domain, projectPath, apacheConfPath string) error {
	vhostConf := fmt.Sprintf(`<VirtualHost *:80>
    ServerName %s
    DocumentRoot "%s"
    <Directory "%s">
        AllowOverride All
        Require all granted
    </Directory>
</VirtualHost>\n`, domain, projectPath, projectPath)
	confFile := filepath.Join(apacheConfPath, domain+".conf")
	return os.WriteFile(confFile, []byte(vhostConf), 0644)
}

func RemoveVHost(domain, apacheConfPath string) error {
	confFile := filepath.Join(apacheConfPath, domain+".conf")
	return os.Remove(confFile)
}
