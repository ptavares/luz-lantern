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

package mapper

import (
	"github.com/rodesousa/lantern/shard"
)

func PatternMatching(key string, value shard.ShardArguments) shard.Shard {
	item := initShard(key)
	item.Args = value
	return item
}

func initShard(key string) shard.Shard {
	switch key {
	case "user":
		return shard.InitUser()
	case "ping":
		return shard.InitPing()
	}
	return shard.InitUnknow()
}