/*
	Copyright © 2021-2023 Macaroni OS Linux
	See AUTHORS and LICENSE for the license details and contributors.
*/

package converter

import (
	"errors"
	"io/ioutil"

	pkg "github.com/macaroni-os/anise/pkg/package"
	st "github.com/macaroni-os/anise/pkg/spectooling"
	compilerspec "github.com/macaroni-os/anise/pkg/v2/compiler/types/specs"

	"gopkg.in/yaml.v2"
)

type AniseCompilationSpecSanitized struct {
	Steps      []string `json:"steps,omitempty" yaml:"steps,omitempty"`
	Env        []string `json:"env,omitempty" yaml:"env,omitempty"`
	Prelude    []string `json:"prelude,omitempty" yaml:"prelude,omitempty"` // Are run inside the image which will be our builder
	Image      string   `json:"image,omitempty" yaml:"image,omitempty"`
	Seed       string   `json:"seed,omitempty" yaml:"seed,omitempty"`
	PackageDir string   `json:"package_dir,omitempty" yaml:"package_dir,omitempty"`

	Retrieve []string `json:"retrieve,omitempty" yaml:"retrieve,omitempty"`

	Unpack   bool     `json:"unpack,omitempty" yaml:"unpack,omitempty"`
	Includes []string `json:"includes,omitempty" yaml:"includes,omitempty"`
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	PackageRequires  []*st.DefaultPackageSanitized `json:"requires,omitempty" yaml:"requires,omitempty"`
	PackageConflicts []*st.DefaultPackageSanitized `json:"conflicts,omitempty" yaml:"conflicts,omitempty"`
	Provides         []*st.DefaultPackageSanitized `json:"provides,omitempty" yaml:"provides,omitempty"`
}

func NewAniseCompilationSpecSanitized(s *compilerspec.CompilationSpec) *AniseCompilationSpecSanitized {
	return &AniseCompilationSpecSanitized{
		Steps:      s.Steps,
		Env:        s.Env,
		Prelude:    s.Prelude,
		Image:      s.Image,
		Seed:       s.Seed,
		PackageDir: s.PackageDir,
		Retrieve:   s.Retrieve,
		Unpack:     s.Unpack,
		Includes:   s.Includes,
		Excludes:   s.Excludes,
	}
}

func NewAniseCompilationSpecSanitizedFromYaml(data []byte) (*AniseCompilationSpecSanitized, error) {
	ans := &AniseCompilationSpecSanitized{}
	if err := yaml.Unmarshal(data, ans); err != nil {
		return nil, err
	}
	return ans, nil
}

func NewAniseCompilationSpecSanitizedFromFile(file string) (*AniseCompilationSpecSanitized, error) {
	if file == "" {
		return nil, errors.New("Invalid file path")
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	ans, err := NewAniseCompilationSpecSanitizedFromYaml(content)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func (s *AniseCompilationSpecSanitized) GetRequires() []*st.DefaultPackageSanitized {
	return s.PackageRequires
}

func (s *AniseCompilationSpecSanitized) GetConflicts() []*st.DefaultPackageSanitized {
	return s.PackageConflicts
}

func (s *AniseCompilationSpecSanitized) Yaml() ([]byte, error) {
	return yaml.Marshal(s)
}

func (s *AniseCompilationSpecSanitized) WriteBuildDefinition(path string) error {
	data, err := s.Yaml()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func (s *AniseCompilationSpecSanitized) Clone() (*AniseCompilationSpecSanitized, error) {
	data, err := s.Yaml()
	if err != nil {
		return nil, err
	}

	return NewAniseCompilationSpecSanitizedFromYaml(data)
}

func (s *AniseCompilationSpecSanitized) AddRequires(pkgs []*pkg.DefaultPackage) {
	for _, pkg := range pkgs {
		p := st.NewDefaultPackageSanitized(pkg)
		s.PackageRequires = append(s.PackageRequires, p)
	}
}

func (s *AniseCompilationSpecSanitized) Requires(pkgs []*pkg.DefaultPackage) {
	s.PackageRequires = []*st.DefaultPackageSanitized{}
	s.AddRequires(pkgs)
}

func (s *AniseCompilationSpecSanitized) AddConflicts(pkgs []*pkg.DefaultPackage) {
	for _, pkg := range pkgs {
		p := st.NewDefaultPackageSanitized(pkg)
		s.PackageConflicts = append(s.PackageConflicts, p)
	}
}

func (s *AniseCompilationSpecSanitized) Conflicts(pkgs []*pkg.DefaultPackage) {
	s.PackageConflicts = []*st.DefaultPackageSanitized{}
	s.AddConflicts(pkgs)
}
