package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const DashDir = "gh-dash"
const ConfigFileName = "config.yml"

type ViewType string

const (
	PRsView    ViewType = "prs"
	IssuesView ViewType = "issues"
)

type SectionConfig struct {
	Title   string
	Filters string
	Limit   *int `yaml:"limit,omitempty"`
}

type PreviewConfig struct {
	Open  bool
	Width int
}

type Defaults struct {
	Preview     PreviewConfig `yaml:"preview"`
	PrsLimit    int           `yaml:"prsLimit"`
	IssuesLimit int           `yaml:"issuesLimit"`
	View        ViewType      `yaml:"view"`
}

type Config struct {
	PRSections     []SectionConfig `yaml:"prSections"`
	IssuesSections []SectionConfig `yaml:"issuesSections"`
	Defaults       Defaults        `yaml:"defaults"`
}

type configError struct {
	configDir string
	parser    ConfigParser
	err       error
}

type ConfigParser struct{}

func (parser ConfigParser) getDefaultConfig() Config {
	return Config{
		Defaults: Defaults{
			Preview: PreviewConfig{
				Open:  true,
				Width: 50,
			},
			PrsLimit:    20,
			IssuesLimit: 20,
			View:        PRsView,
		},
		PRSections: []SectionConfig{
			{
				Title:   "Mis Pr`s",
				Filters: "is:open author:@me",
			},
						{
				Title:   "TciManager",
				filters: is:open -author:@me repo:vnikiet/promanager
			},
			{
				Title:   "Needs My Review",
				Filters: "is:open review-requested:@me",
			},
			{
				Title:   "Subscribed",
				Filters: "is:open -author:@me repo:cli/cli repo:dlvhdr/gh-dash",
			},
		},
		IssuesSections: []SectionConfig{
			{
				Title:   "My Issues",
				Filters: "is:open author:@me",
			},
			{
				Title:   "Assigned",
				Filters: "is:open assignee:@me",
			},
			{
				Title:   "Subscribed",
				Filters: "is:open -author:@me repo:cli/cli repo:dlvhdr/gh-dash",
			},
		},
	}
}

func (parser ConfigParser) getDefaultConfigYamlContents() string {
	defaultConfig := parser.getDefaultConfig()
	yaml, _ := yaml.Marshal(defaultConfig)

	return string(yaml)
}

func (e configError) Error() string {
	return fmt.Sprintf(
		`Couldn't find a config.yml configuration file.
Create one under: %s

Example of a config.yml file:
%s

For more info, go to https://github.com/dlvhdr/gh-dash
press q to exit.

Original error: %v`,
		path.Join(e.configDir, DashDir, ConfigFileName),
		string(e.parser.getDefaultConfigYamlContents()),
		e.err,
	)
}

func (parser ConfigParser) writeDefaultConfigContents(newConfigFile *os.File) error {
	_, err := newConfigFile.WriteString(parser.getDefaultConfigYamlContents())

	if err != nil {
		return err
	}

	return nil
}

func (parser ConfigParser) createConfigFileIfMissing(configFilePath string) error {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		newConfigFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}

		defer newConfigFile.Close()
		return parser.writeDefaultConfigContents(newConfigFile)
	}

	return nil
}

func (parser ConfigParser) getConfigFileOrCreateIfMissing() (*string, error) {
	var err error
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir, err = os.UserConfigDir()
		if err != nil {
			return nil, configError{parser: parser, configDir: configDir, err: err}
		}
	}

	prsConfigDir := filepath.Join(configDir, DashDir)
	err = os.MkdirAll(prsConfigDir, os.ModePerm)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	configFilePath := filepath.Join(prsConfigDir, ConfigFileName)
	err = parser.createConfigFileIfMissing(configFilePath)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	return &configFilePath, nil
}

type parsingError struct {
	err error
}

func (e parsingError) Error() string {
	return fmt.Sprintf("failed parsing config.yml: %v", e.err)
}

func (parser ConfigParser) readConfigFile(path string) (Config, error) {
	config := parser.getDefaultConfig()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal([]byte(data), &config)
	return config, err
}

func initParser() ConfigParser {
	return ConfigParser{}
}

func ParseConfig() (Config, error) {
	parser := initParser()
	var config Config
	var err error
	configFilePath, err := parser.getConfigFileOrCreateIfMissing()
	if err != nil {
		return config, parsingError{err: err}
	}

	config, err = parser.readConfigFile(*configFilePath)
	if err != nil {
		return config, parsingError{err: err}
	}

	return config, nil
}
