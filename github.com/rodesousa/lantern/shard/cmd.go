// Copyright © 2016 Roberto De Sousa (https://github.com/rodesousa) / Patrick Tavares (https://github.com/ptavares)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package shard provides ...
package shard

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getExpected(shard *Shard) bool {
	if val, ok := shard.Args["expected"]; ok {
		return val.(bool)
	}
	return true
}

func (s *Shard) Cmd() bool {
	out, status, error, err := s.exeCmd()
	if getExpected(s) != status {
		if error != nil {
			s.Status.Err = fmt.Sprintf("\n %s", err)
			s.Status.Check = false
		}

		if s.Checked.Enabled {
			if strings.Contains(string(out[:]), ValueChecked) != false {
				s.Status.Err = fmt.Sprintf("\n %s", out)
				s.Status.Check = false
			}
		}
	}
	return s.Status.Check
}

func (s Shard) exeCmd() (string, bool, error, string) {
	c := exec.Command(s.Command, s.CommandArguments...)

	err := c.Run()
	stdout := fmt.Sprint(os.Stdout)

	if err != nil {
		return stdout, false, err, fmt.Sprint(os.Stderr)
	} else {
		return stdout, true, nil, ""
	}
}
