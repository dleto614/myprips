# myprips
I created this tool from the ground up due to not being satisfied with the way the original prips worked and it had limitations.
I added a few more basic command line arguments along with STDIN input.

You can compile this program by running:

```
git clone https://github.com/dleto614/myprips
cd myprips && go build
```

--------------

Usage:

```
Usage of ./prips:
  -i string
        Specify input file with ip ranges.
  -o string
        Specify the output file to write results into.
  -r string
        Specify ip range.
```

--------------

STDIN:

```
$ echo "192.168.1.0/24" | ./prips
192.168.1.0
192.168.1.1
192.168.1.2
192.168.1.3
192.168.1.4
192.168.1.5
192.168.1.6
...
```

---

```
$ echo "192.168.1.0/24" | ./prips -o test-stdin.txt 
[*] Reading stdin
[*] Checking: 192.168.1.0/24 
```

----------

Input file:

```
$ ./prips -i test-cidr.txt -o test-cidr-file.txt
[*] Reading file: test-cidr.txt 
[*] Checking: 192.168.1.1/24 
[*] Checking: 142.250.0.0/15
... 
```
---

```
$ ./prips -i test-cidr.txt
...
205.251.243.248
205.251.243.249
205.251.243.250
205.251.243.251
205.251.243.252
205.251.243.253
205.251.243.254
205.251.243.255
```

----------

Range:

```
$ ./prips -r 167.220.226.0/23 -o test-range.txt
[*] Checking: 167.220.226.0/23 
```

---

```
./prips -r 167.220.226.0/23
167.220.226.0
167.220.226.1
167.220.226.2
167.220.226.3
167.220.226.4
167.220.226.5
...
```

