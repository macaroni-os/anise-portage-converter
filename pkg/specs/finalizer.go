/*
Copyright (C) 2020-2021  Daniele Rondina <geaaru@sabayonlinux.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/
package specs

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (f *Finalizer) Yaml() ([]byte, error) {
	return yaml.Marshal(f)
}

func (f *Finalizer) WriteFinalize(path string) error {
	data, err := f.Yaml()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func (f *Finalizer) IsValid() bool {
	if len(f.Install) > 0 {
		return true
	}
	return false
}
