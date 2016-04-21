/*
 * Copyright 2016 ThoughtWorks, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package agent

import (
	"github.com/gocd-contrib/gocd-golang-agent/protocol"
)

func CommandCond(s *BuildSession, cmd *protocol.BuildCommand) error {
	for i := 0; i < len(cmd.SubCommands); i += 2 {
		if i == len(cmd.SubCommands)-1 {
			// else branch
			return s.process(cmd.SubCommands[i])
		}
		test := cmd.SubCommands[i]
		action := cmd.SubCommands[i+1]
		_, err := s.processTestCommand(test)
		if err == nil {
			return s.process(action)
		}
	}
	return nil
}
