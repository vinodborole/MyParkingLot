## Build

Note: GOPATH for MyParkingLot is set automatically by the build.sh script 

```
$ cd MyParkingLot/scripts

$ sh build.sh

```
## Execute

```
$ ./parking_lot file_input.txt
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No.	Registration No.	Colour
1	KA-01-HH-1234	White
2	KA-01-HH-9999	White
3	KA-01-BB-0001	Black
5	KA-01-HH-2701	Blue
6	KA-01-HH-3141	Black
Allocated slot number: 4
sorry, parking lot is full
KA-01-HH-1234,KA-01-HH-9999,KA-01-P-333
not found
6
not found
```