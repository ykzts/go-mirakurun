// Copyright 2018 Yamagishi Kazutoshi
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mirakurun_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ykzts/go-mirakurun/mirakurun"
)

func ExampleNewClient() {
	c := mirakurun.NewClient()
	c.BaseURL, _ = url.Parse("http://192.168.0.5:40772/api/")

	programs, _, err := c.GetPrograms(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Program count: ", len(programs))
}

func ExampleClient_GetChannels() {
	c := mirakurun.NewClient()

	channels, _, err := c.GetChannels(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, channel := range channels {
		fmt.Printf("%s (%s): %s\n", channel.Channel, channel.Type, channel.Name)
	}
}

func ExampleClient_GetChannelsByType() {
	c := mirakurun.NewClient()

	channels, _, err := c.GetChannelsByType(context.Background(), "GR", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, channel := range channels {
		fmt.Printf("%s (%s): %s\n", channel.Channel, channel.Type, channel.Name)
	}
}

func ExampleClient_GetChannel() {
	c := mirakurun.NewClient()

	channel, _, err := c.GetChannel(context.Background(), "GR", "16")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s (%s): %s\n", channel.Channel, channel.Type, channel.Name)
}

func ExampleClient_GetServicesByChannel() {
	c := mirakurun.NewClient()

	services, _, err := c.GetServicesByChannel(context.Background(), "GR", "16")
	if err != nil {
		log.Fatal(err)
	}

	for _, service := range services {
		fmt.Printf("%d: %s\n", service.ID, service.Name)
	}
}

func ExampleClient_GetServiceByChannel() {
	c := mirakurun.NewClient()

	service, _, err := c.GetServiceByChannel(context.Background(), "GR", "16", 3239123608)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", service.ID, service.Name)
}

func ExampleClient_GetServiceStreamByChannel() {
	filename := fmt.Sprintf("/tmp/stream-%d.ts", time.Now().Unix())

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := mirakurun.NewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, _, err := c.GetServiceStreamByChannel(ctx, "GR", "16", 3239123608, true)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	fmt.Println("output: ", filename)
	io.Copy(file, stream)
}

func ExampleClient_GetPrograms() {
	c := mirakurun.NewClient()

	opt := &mirakurun.ProgramsListOptions{ServiceID: 23608}
	programs, _, err := c.GetPrograms(context.Background(), opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, program := range programs {
		fmt.Printf("%d: %s (%v)\n", program.ID, program.Name, program.StartAt)
	}
}

func ExampleClient_GetProgram() {
	c := mirakurun.NewClient()

	program, _, err := c.GetProgram(context.Background(), 323912360802956)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s (%v)\n", program.ID, program.Name, program.StartAt)
	fmt.Println(program.Description)

	if program.Extended != nil {
		fmt.Println("")
		for key, value := range program.Extended {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}

func ExampleClient_GetProgramStream() {
	filename := fmt.Sprintf("/tmp/stream-%d.ts", time.Now().Unix())

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := mirakurun.NewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, _, err := c.GetProgramStream(ctx, 323912360802956, true)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	fmt.Println("output: ", filename)
	io.Copy(file, stream)
}

func ExampleClient_GetServices() {
	c := mirakurun.NewClient()

	services, _, err := c.GetServices(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, service := range services {
		fmt.Printf("%d: %s\n", service.ID, service.Name)
	}
}

func ExampleClient_GetService() {
	c := mirakurun.NewClient()

	service, _, err := c.GetService(context.Background(), 3239123608)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s\n", service.ID, service.Name)
}

func ExampleClient_GetLogoImage() {
	filename := fmt.Sprintf("/tmp/logo-%d.png", time.Now().Unix())

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := mirakurun.NewClient()

	stream, _, err := c.GetLogoImage(context.Background(), 3239123608)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	fmt.Println("output: ", filename)
	io.Copy(file, stream)
}

func ExampleClient_GetServiceStream() {
	filename := fmt.Sprintf("/tmp/stream-%d.ts", time.Now().Unix())

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := mirakurun.NewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, _, err := c.GetServiceStream(ctx, 3239123608, true)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	fmt.Println("output: ", filename)
	io.Copy(file, stream)
}

func ExampleClient_GetTuners() {
	c := mirakurun.NewClient()

	tuners, _, err := c.GetTuners(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, tuner := range tuners {
		fmt.Printf("%d: %s (%s)\n", tuner.Index, tuner.Name, strings.Join(tuner.Types, ", "))
	}
}

func ExampleClient_GetTuner() {
	c := mirakurun.NewClient()

	tuner, _, err := c.GetTuner(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d: %s (%s)\n", tuner.Index, tuner.Name, strings.Join(tuner.Types, ", "))
}

func ExampleClient_GetTunerProcess() {
	c := mirakurun.NewClient()

	process, _, err := c.GetTunerProcess(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PID: ", process.PID)
}

func ExampleClient_KillTunerProcess() {
	c := mirakurun.NewClient()

	process, _, err := c.KillTunerProcess(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PID: ", process.PID)
}

func ExampleClient_GetEvents() {
	c := mirakurun.NewClient()

	events, _, err := c.GetEvents(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Printf("%s: %s\n", event.Resource, event.Type)
	}
}

func ExampleClient_GetEventsStream() {
	c := mirakurun.NewClient()

	stream, _, err := c.GetEventsStream(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	io.Copy(os.Stdout, stream)
}

func ExampleClient_GetLog() {
	c := mirakurun.NewClient()

	buf, _, err := c.GetLog(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf)
}

func ExampleClient_GetLogStream() {
	c := mirakurun.NewClient()

	stream, _, err := c.GetLogStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	io.Copy(os.Stdout, stream)
}

func ExampleClient_CheckVersion() {
	c := mirakurun.NewClient()

	version, _, err := c.CheckVersion(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current: " + version.Current)
	fmt.Println("Latest: " + version.Latest)
}

func ExampleClient_UpdateVersion() {
	c := mirakurun.NewClient()

	stream, _, err := c.UpdateVersion(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	io.Copy(os.Stdout, stream)
}

func ExampleClient_GetStatus() {
	c := mirakurun.NewClient()

	status, _, err := c.GetStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version: " + status.Version)
}

func ExampleClient_Restart() {
	c := mirakurun.NewClient()

	restart, _, err := c.Restart(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PID: ", restart.PID)
}
