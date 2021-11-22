/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"mosn.io/layotto/sdk/go-sdk/client"
)

const (
	topic = "in-memory"
)

func main() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	err = testPublish(cli)
	if err != nil {
		panic(err)
	}
}

func testPublish(cli client.Client) error {
	return cli.PublishEvent(context.Background(), "in-memory", topic, []byte("hello in-memory pubsub"))
}

func testSubscribe(cli client.Client) error {
	return nil
}
