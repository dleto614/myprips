package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"slices"
	"strings"
)

func main() {

	var file *string
	var output *string
	var cidr *string

	var ip net.IP

	var ProcessedSlice []string
	var data []string

	file = flag.String("i", "", "Specify input file with ip ranges.")
	output = flag.String("o", "", "Specify the output file to write results into.")
	cidr = flag.String("r", "", "Specify ip range.")

	flag.Parse()

	results := ChkStdin(os.Stdin)

	if len(results) > 0 {
		if ChkFlag("o") {
			fmt.Println("[*] Reading stdin")
		}

		data = append(data, strings.TrimSpace(results))
	} else if ChkFlag("i") {
		if ChkFlag("o") {
			fmt.Printf("[*] Reading file: %s \n", *file)
		}

		data = ReadFile(file)
	} else if ChkFlag("r") {
		data = append(data, *cidr)
	} else {
		usage()
		os.Exit(1)
	}

	for _, DataString := range data {

		if ChkFlag("o") {
			fmt.Printf("[*] Checking: %s \n", DataString)
		}

		Ipv4Addr, Ipv4Net, err := net.ParseCIDR(DataString)

		if err != nil {
			log.Fatal(err)
		}

		for ip = Ipv4Addr.Mask(Ipv4Net.Mask); Ipv4Net.Contains(ip); IncrementIP(ip) {

			if ChkFlag("o") && !slices.Contains(ProcessedSlice, ip.String()) {
				FileWrite(ip.String(), *output)

				// Add the processed string to slice to prevent writing duplicates
				ProcessedSlice = append(ProcessedSlice, ip.String())
			}

			// Write to stdin if no output file arg was passed
			if !ChkFlag("o") {
				fmt.Println(ip)
			}
		}
	}

	/*
		Original test code block. Leaving this here as documentation.

		ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
		if err != nil {
			log.Fatal(err)
		}

		for ip := ipv4Addr.Mask(ipv4Net.Mask); ipv4Net.Contains(ip); inc(ip) {
			fmt.Println(ip)
		}
	*/
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s: \n", os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
}

func ChkFlag(name string) bool {
	found := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}

func ChkStdin(stdin *os.File) string {

	var results string

	stat, err := stdin.Stat()

	if err != nil {
		log.Fatal(err)
	}

	// I don't know why this works and too tired to look into it more, but not just checking the size doesn't and it does work in main.
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := io.ReadAll(stdin)

		if err != nil {
			log.Fatal(err)
		}

		results = string(bytes)
	}

	return results
}

func IncrementIP(ip net.IP) {
	for num := len(ip) - 1; num >= 0; num-- {
		ip[num]++
		if ip[num] > 0 {
			break
		}
	}
}

func ReadFile(file *string) []string {

	var data []string

	fd, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	fd.Close()
	return data
}

func FileWrite(data string, data_file string) {

	file, err := os.OpenFile(data_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Create and open file to append into

	if err != nil {
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil { // Write data to file
		return
	}

}
