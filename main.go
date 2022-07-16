package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type stopConfig struct {
	ClusterName  string   `json:"cluster_name"`
	ServiceNames []string `json:"service_names"`
}

var (
	help       = flag.Bool("help", false, "output usage")
	configFile = flag.String("config_file", "config.json", "stop configuration file path")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	f, err := os.Open(*configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	d := json.NewDecoder(f)

	var esn stopConfig
	if err := d.Decode(&esn); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := ecs.NewFromConfig(cfg)
	for _, v := range esn.ServiceNames {
		_, err := client.UpdateService(context.Background(), &ecs.UpdateServiceInput{
			DesiredCount: aws.Int32(0),
			Cluster:      &esn.ClusterName,
			Service:      &v,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
