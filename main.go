package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gcp-calculator-mcp/calculator"
	"gcp-calculator-mcp/mcp"
)

func main() {
	if len(os.Args) > 1 {
		subcommand := os.Args[1]
		switch subcommand {
		case "compute":
			runComputeCLI(os.Args[2:])
			return
		case "cloud-run", "run":
			runCloudRunCLI(os.Args[2:])
			return
		case "storage":
			runStorageCLI(os.Args[2:])
			return
		case "gke", "kubernetes":
			runGKECLI(os.Args[2:])
			return
		case "cloud-sql", "sql":
			runCloudSQLCLI(os.Args[2:])
			return
		case "bigquery", "bq":
			runBigQueryCLI(os.Args[2:])
			return
		case "service":
			runUniversalCLI(os.Args[2:])
			return
		case "catalog", "search":
			runCatalogCLI(os.Args[2:])
			return
		case "estimate":
			runEstimateCLI(os.Args[2:])
			return
		case "pricing", "policy":
			runPolicyCLI()
			return
		case "server":
			runMCPServer()
			return
		case "-h", "--help", "help":
			printHelp()
			return
		}
	}

	runMCPServer()
}

func runMCPServer() {
	server := mcp.NewServer(os.Stdin, os.Stdout)
	if err := server.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "MCP server error: %v\n", err)
		os.Exit(1)
	}
}

func runPolicyCLI() {
	policy := calculator.GetActivePolicy()
	printJSON(policy)
}

func runComputeCLI(args []string) {
	fs := flag.NewFlagSet("compute", flag.ExitOnError)
	machine := fs.String("machine", "e2-micro", "Machine type")
	hours := fs.Float64("hours", 730, "Operating hours/month")
	spot := fs.Bool("spot", false, "Use Spot VM")
	region := fs.String("region", "us-central1", "GCP Region (e.g. us-central1, asia-northeast3)")
	commitment := fs.String("commitment", "none", "Commitment Discount (none, 1-year, 3-year)")
	gpuType := fs.String("gpu-type", "", "GPU model (e.g. nvidia-tesla-t4, nvidia-l4, nvidia-tesla-a100-80gb, nvidia-h100-80gb)")
	gpuCount := fs.Int("gpu-count", 0, "Number of GPUs")
	tpuType := fs.String("tpu-type", "", "TPU model (e.g. tpu-v2-8, tpu-v3-8, tpu-v4-pod-8, tpu-v5litepod-8, tpu-v5p-8)")
	tpuCount := fs.Int("tpu-count", 0, "Number of TPUs")
	storage := fs.Float64("storage", 20, "Disk size in GB")
	storageType := fs.String("storage-type", "standard", "Disk type")
	_ = fs.Parse(args)

	input := calculator.ComputeEngineInput{
		MachineType:   *machine,
		Region:        *region,
		Commitment:    *commitment,
		HoursPerMonth: *hours,
		IsSpot:        *spot,
		GPUType:       *gpuType,
		GPUCount:      *gpuCount,
		TPUType:       *tpuType,
		TPUCount:      *tpuCount,
		StorageGB:     *storage,
		StorageType:   *storageType,
	}

	res := calculator.CalculateComputeEngine(input)
	printJSON(res)
}

func runCloudRunCLI(args []string) {
	fs := flag.NewFlagSet("cloud-run", flag.ExitOnError)
	requests := fs.Float64("requests", 1000000, "Requests/month")
	cpu := fs.Float64("cpu", 1.0, "vCPUs")
	mem := fs.Float64("memory", 0.5, "Memory GB")
	duration := fs.Float64("duration", 200, "Duration ms")
	concurrency := fs.Float64("concurrency", 80, "Concurrency")
	_ = fs.Parse(args)

	input := calculator.CloudRunInput{
		RequestsPerMonth: *requests,
		CPUCores:         *cpu,
		MemoryGB:         *mem,
		AvgDurationMS:    *duration,
		Concurrency:      *concurrency,
	}

	res := calculator.CalculateCloudRun(input)
	printJSON(res)
}

func runStorageCLI(args []string) {
	fs := flag.NewFlagSet("storage", flag.ExitOnError)
	storage := fs.Float64("gb", 50, "Storage GB")
	class := fs.String("class", "standard", "Storage class")
	egress := fs.Float64("egress", 5, "Network egress GB")
	_ = fs.Parse(args)

	input := calculator.StorageInput{
		StorageGB:       *storage,
		StorageClass:    *class,
		NetworkEgressGB: *egress,
	}

	res := calculator.CalculateStorage(input)
	printJSON(res)
}

func runGKECLI(args []string) {
	fs := flag.NewFlagSet("gke", flag.ExitOnError)
	mode := fs.String("mode", "autopilot", "Mode: autopilot or standard")
	nodes := fs.Int("nodes", 3, "Node count")
	vcpus := fs.Float64("vcpus", 2.0, "vCPUs per node")
	mem := fs.Float64("memory", 4.0, "Memory GB per node")
	_ = fs.Parse(args)

	input := calculator.GKEInput{
		Mode:         *mode,
		NodeCount:    *nodes,
		VCPUsPerNode: *vcpus,
		MemGBPerNode: *mem,
	}

	res := calculator.CalculateGKE(input)
	printJSON(res)
}

func runCloudSQLCLI(args []string) {
	fs := flag.NewFlagSet("cloud-sql", flag.ExitOnError)
	instance := fs.String("instance", "db-f1-micro", "Instance type")
	storage := fs.Float64("gb", 10, "Storage GB")
	ha := fs.Bool("ha", false, "High availability mode")
	_ = fs.Parse(args)

	input := calculator.CloudSQLInput{
		InstanceType: *instance,
		StorageGB:    *storage,
		HighAvail:    *ha,
	}

	res := calculator.CalculateCloudSQL(input)
	printJSON(res)
}

func runBigQueryCLI(args []string) {
	fs := flag.NewFlagSet("bigquery", flag.ExitOnError)
	queryTB := fs.Float64("query-tb", 2.0, "Scanned query volume in TB")
	storageGB := fs.Float64("storage-gb", 50.0, "Active storage in GB")
	_ = fs.Parse(args)

	input := calculator.BigQueryInput{
		QueryTBScanned:  *queryTB,
		ActiveStorageGB: *storageGB,
	}

	res := calculator.CalculateBigQuery(input)
	printJSON(res)
}

func runCatalogCLI(args []string) {
	if len(args) > 0 {
		query := args[0]
		results := calculator.SearchCatalog(query)
		printJSON(map[string]interface{}{
			"query":   query,
			"count":   len(results),
			"results": results,
		})
		return
	}

	catalog := calculator.GetFullCatalog()
	printJSON(map[string]interface{}{
		"total_count": len(catalog),
		"catalog":     catalog,
	})
}

func runUniversalCLI(args []string) {
	if len(args) > 1 {
		svcName := args[0]
		var params map[string]interface{}
		_ = json.Unmarshal([]byte(args[1]), &params)
		input := calculator.UniversalInput{
			ServiceName: svcName,
			Params:      params,
		}
		res := calculator.CalculateService(input)
		printJSON(res)
		return
	}
	fmt.Println("Usage: gcp-calculator-mcp service <SERVICE_NAME> '<JSON_PARAMS>'")
}

func runEstimateCLI(args []string) {
	if len(args) > 0 {
		var input calculator.EstimateInput
		err := json.Unmarshal([]byte(args[0]), &input)
		if err == nil {
			res := calculator.EstimateInfrastructure(input)
			printJSON(res)
			return
		}
	}
	fmt.Println("Usage: gcp-calculator-mcp estimate '<JSON_SPEC>'")
}

func printJSON(val interface{}) {
	bytes, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

func printHelp() {
	fmt.Println(`GCP Pricing Calculator MCP Server & CLI Tool (Golang - All Products)

Usage:
  gcp-calculator-mcp [command] [flags]

Commands:
  server      Run as MCP stdio server (default)
  compute     Calculate Compute Engine costs (VMs, Spot, GPUs, Disks)
  cloud-run   Calculate Cloud Run costs
  storage     Calculate Cloud Storage costs
  gke         Calculate Google Kubernetes Engine costs
  cloud-sql   Calculate Cloud SQL costs
  bigquery    Calculate BigQuery costs
  service     Calculate cost for any GCP service dynamically
  estimate    Calculate full infrastructure cost from JSON spec
  pricing     Display current active pricing policy rates

Examples:
  gcp-calculator-mcp compute -machine e2-micro -spot
  gcp-calculator-mcp compute -machine n1-standard-4 -gpu-type nvidia-tesla-t4 -gpu-count 1 -spot
  gcp-calculator-mcp gke -mode autopilot -nodes 3 -vcpus 2 -memory 4
  gcp-calculator-mcp cloud-sql -instance db-f1-micro -gb 20
  gcp-calculator-mcp bigquery -query-tb 5 -storage-gb 200
  gcp-calculator-mcp pricing
`)
}
