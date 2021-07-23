/*
Copyright © 2021 aoyouer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"behaviorlog-analyzer/cmd"
	"behaviorlog-analyzer/server"
	"log"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			if err, ok := r.(error); ok {
				log.Println("出现错误", err.Error())
			} else {
				log.Printf("异常 %v", r)
			}
		}
	}()
	go server.Start("127.0.0.1", "8080")
	cmd.Execute()
}

