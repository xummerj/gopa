/*
Copyright 2016 Medcl (m AT medcl.net)

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

package logging

import (
	log "github.com/cihub/seelog"
	. "github.com/medcl/gopa/core/env"
	"strings"
)

func SetInitLogging(env *Env, logLevel string) {

	logLevel = strings.ToLower(logLevel)

	testConfig := `
	<seelog  type="sync" minlevel="`
	testConfig = testConfig + logLevel
	testConfig = testConfig + `">
		<outputs formatid="main">
			<filter levels="error">
				<file path="./log/gopa.log"/>
			</filter>
			<console formatid="main" />
		</outputs>
		<formats>
			<format id="main" format="[%Date(01-02) %Time] [%LEV] [%File:%Line,%FuncShort] %Msg%n"/>
		</formats>
	</seelog>`
	ReplaceConfig(env, testConfig)
}

func SetLogging(env *Env, logLevel string, logFile string) {

	logLevel = strings.ToLower(logLevel)

	testConfig := `
	<seelog  type="sync" minlevel="`
	testConfig = testConfig + logLevel
	testConfig = testConfig + `">
		<outputs formatid="main">
			<console formatid="main"/>
			<filter levels="` + logLevel + `">
				<file path="` + logFile + `"/>
			</filter>
			 <rollingfile formatid="main" type="size" filename="` + logFile + `" maxsize="10000000000" maxrolls="5" />
		</outputs>
		<formats>
			<format id="main" format="[%Date(01-02) %Time] [%LEV] [%File:%Line,%FuncShort] %Msg%n"/>
		</formats>
	</seelog>`
	ReplaceConfig(env, testConfig)
}

func ReplaceConfig(env *Env, cfg string) {
	logger, err := log.LoggerFromConfigAsString(cfg)
	if err != nil {
		log.Error("replace config error,", err)
	}
	err = log.ReplaceLogger(logger)
	if err != nil {
		log.Error("replace config error,", err)
	}
	if env != nil && env.RuntimeConfig != nil {
		env.RuntimeConfig.LoggingConfig = cfg
	}
	log.Info("config success replaced")
}

func GetLoggingConfig(env *Env) string {
	return env.RuntimeConfig.LoggingConfig
}

func Flush() {
	log.Flush()
}
