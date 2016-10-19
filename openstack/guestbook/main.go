/*
Copyright 2014 The Kubernetes Authors.

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
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var (
	messages []string
)

func ListRangeHandler(rw http.ResponseWriter, req *http.Request) {
	membersJSON := HandleError(json.MarshalIndent(messages, "", "  ")).([]byte)
	rw.Write(membersJSON)
}

func ListPushHandler(rw http.ResponseWriter, req *http.Request) {
	value := mux.Vars(req)["value"]
	if alreadyExists(value) {
		fmt.Printf("Error: %s already exists\n", value)
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		messages = append(messages, value)
		fmt.Printf("Added: %s \n", value)
		ListRangeHandler(rw, req)
	}
}

func alreadyExists(value string) bool {
	for _, message := range messages {
		if message == value {
			return true
		}
	}
	return false
}

func InfoHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(nil)
}

func EnvHandler(rw http.ResponseWriter, req *http.Request) {
	environment := make(map[string]string)
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]
		val := strings.Join(splits[1:], "=")
		environment[key] = val
	}
	envJSON := HandleError(json.MarshalIndent(environment, "", "  ")).([]byte)
	rw.Write(envJSON)
}

func CpuShortHandler(rw http.ResponseWriter, req *http.Request) {
	x := 0.0001
	t := (rand.Intn(8) + 2) * 100000
	for i := 0; i <= t; i++ {
		x += math.Jn(100, x)
	}
	rw.Write([]byte("."))
}

func CpuLongHandler(rw http.ResponseWriter, req *http.Request) {
	t := time.Millisecond * 250
	for i := 0; i <= 20; i++ {
		CpuShortHandler(rw, req)
		fmt.Printf("Round %v of cpu usage finished\n",i)
		time.Sleep(t)
	}
	rw.Write([]byte("All done"))
}

func HealthHandler(rw http.ResponseWriter, req *http.Request){
	rw.Write([]byte("ok"))
}

func HandleError(result interface{}, err error) (r interface{}) {
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	rand.Seed(42)
	messages = make([]string, 10)

	r := mux.NewRouter()
	r.Path("/list").Methods("GET").HandlerFunc(ListRangeHandler)
	r.Path("/push/{value}").Methods("GET").HandlerFunc(ListPushHandler)
	r.Path("/info").Methods("GET").HandlerFunc(InfoHandler)
	r.Path("/env").Methods("GET").HandlerFunc(EnvHandler)
	r.Path("/cpu-short").Methods("GET").HandlerFunc(CpuShortHandler)
	r.Path("/cpu-long").Methods("GET").HandlerFunc(CpuLongHandler)
	r.Path("/health").Methods("GET").HandlerFunc(HealthHandler)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3010")
}
