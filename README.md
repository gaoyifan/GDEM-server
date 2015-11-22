# GDEM-server
The height data server for GDEM

# Install via docker

	docker run --restart=always --name=redis -v /path/to/redis:/data -d redis:3
    docker run --restart=always --name=gdemd -v /path/to/map:/srv/map --link redis:redis -d -p 8000:8000 gaoyfian/gdem-server

# Download Map Data

You can downlaod the GDEM v2 Map Data(global data) which is converted to png format:

    http://pan.baidu.com/s/1pJBT1qz

You can just downlaod the data you need.

# API Format

	http://hostname/i/j/zoom/size/
	
* i: abscissa
* j: ordinate
* zoom: Zoom levels ([Reference](http://wiki.openstreetmap.org/wiki/Slippy_map_tilenames))
* size: The resolution of the map block. 2^size * 2^size


# Public Server

You can just use public server via the following URL:

	http://gdem.yfgao.com/

Testing the connection by shell command:

	$ curl http://gdem.yfgao.com/27054/13441/15/4  | hexdump

it will return the following line:

	  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
	                                 Dload  Upload   Total   Spent    Left  Speed
	100   512  100   512    0     0   3768      0 --:--:-- --:--:-- --:--:--  3792
	0000000 1c 00 1c 00 1d 00 38 00 3b 00 2b 00 21 00 2b 00
	0000010 32 00 24 00 22 00 1b 00 18 00 15 00 1b 00 1d 00
	0000020 24 00 19 00 21 00 36 00 34 00 2a 00 33 00 3c 00
	0000030 36 00 2f 00 2b 00 1f 00 1b 00 1b 00 17 00 18 00
	0000040 16 00 13 00 1c 00 2b 00 35 00 3c 00 4a 00 4e 00
	0000050 42 00 3d 00 30 00 26 00 20 00 1e 00 22 00 12 00
	0000060 11 00 14 00 20 00 29 00 36 00 45 00 49 00 4f 00
	0000070 4d 00 43 00 34 00 20 00 24 00 2a 00 2a 00 24 00
    0000080 18 00 10 00 1c 00 27 00 30 00 3b 00 3e 00 4a 00
    0000090 50 00 46 00 30 00 1c 00 28 00 34 00 31 00 23 00
    00000a0 11 00 0f 00 14 00 1d 00 1e 00 1d 00 2d 00 3f 00
    00000b0 40 00 39 00 25 00 22 00 37 00 3c 00 3a 00 30 00
    00000c0 0f 00 10 00 18 00 1a 00 18 00 1b 00 2b 00 39 00
    00000d0 2f 00 2a 00 37 00 3a 00 46 00 4d 00 4e 00 4d 00
    00000e0 1b 00 0f 00 10 00 15 00 1e 00 2e 00 47 00 48 00
    00000f0 3b 00 43 00 50 00 4a 00 4f 00 4a 00 40 00 44 00
    0000100 1d 00 19 00 14 00 1a 00 29 00 31 00 41 00 58 00
    0000110 52 00 47 00 47 00 55 00 50 00 42 00 3e 00 2e 00
    0000120 0d 00 1b 00 1f 00 1a 00 29 00 3d 00 46 00 5b 00
    0000130 4f 00 3e 00 3e 00 47 00 45 00 30 00 2f 00 18 00
    0000140 0d 00 10 00 1a 00 1f 00 27 00 3d 00 55 00 65 00
    0000150 5a 00 3d 00 32 00 29 00 1f 00 26 00 1a 00 19 00
    0000160 10 00 19 00 19 00 25 00 29 00 38 00 4c 00 5a 00
    0000170 5e 00 3c 00 30 00 28 00 25 00 1d 00 1e 00 27 00
    0000180 10 00 18 00 36 00 35 00 2c 00 31 00 3d 00 55 00
    0000190 5c 00 46 00 3b 00 2e 00 2e 00 24 00 24 00 2f 00
    00001a0 14 00 1b 00 34 00 2d 00 21 00 2c 00 39 00 46 00
    00001b0 4e 00 50 00 3c 00 2a 00 26 00 27 00 28 00 36 00
    00001c0 0c 00 0e 00 1d 00 1d 00 19 00 1a 00 26 00 3d 00
    00001d0 45 00 2d 00 1d 00 1f 00 22 00 25 00 2d 00 31 00
    00001e0 1b 00 13 00 19 00 1d 00 17 00 1a 00 1d 00 28 00
    00001f0 29 00 1c 00 1c 00 2a 00 2c 00 30 00 24 00 20 00
    0000200
