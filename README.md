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
* size: The resolution of the map block. (2^size+1) * (2^size+1)


# Public Server

You can just use public server via the following URL:

	http://gdem.yfgao.com/

Testing the connection by shell command:

	$ curl http://gdem.yfgao.com/27054/13441/15/4  | hexdump

it will return the following line:

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   578  100   578    0     0    190      0  0:00:03  0:00:03 --:--:--   190
0000000 1d 00 1c 00 1d 00 38 00 3b 00 2b 00 21 00 2b 00
0000010 32 00 24 00 22 00 1b 00 18 00 15 00 1b 00 20 00
0000020 10 00 20 00 19 00 18 00 36 00 34 00 2a 00 33 00
0000030 3c 00 36 00 2f 00 2b 00 1f 00 1b 00 1a 00 17 00
0000040 1a 00 13 00 19 00 13 00 17 00 2b 00 35 00 3c 00
0000050 4a 00 4e 00 42 00 3d 00 30 00 26 00 20 00 1e 00
0000060 22 00 1e 00 15 00 14 00 14 00 19 00 29 00 36 00
0000070 45 00 49 00 4f 00 4d 00 43 00 34 00 20 00 24 00
0000080 2a 00 2a 00 24 00 24 00 18 00 0e 00 14 00 29 00
0000090 25 00 32 00 33 00 46 00 4c 00 42 00 27 00 1c 00
00000a0 27 00 31 00 32 00 28 00 21 00 18 00 12 00 13 00
00000b0 1c 00 1d 00 1c 00 27 00 38 00 39 00 32 00 26 00
00000c0 34 00 3d 00 49 00 49 00 3d 00 2e 00 11 00 10 00
00000d0 15 00 1a 00 18 00 1b 00 2b 00 39 00 2f 00 2a 00
00000e0 37 00 3a 00 46 00 54 00 4e 00 52 00 35 00 1a 00
00000f0 0f 00 0e 00 15 00 1e 00 2e 00 47 00 48 00 3b 00
0000100 43 00 50 00 4a 00 4f 00 52 00 40 00 41 00 34 00
0000110 1a 00 19 00 1b 00 1a 00 29 00 31 00 41 00 58 00
0000120 52 00 47 00 47 00 55 00 50 00 48 00 3e 00 32 00
0000130 1f 00 10 00 1b 00 21 00 1a 00 29 00 3d 00 46 00
0000140 5b 00 4f 00 3e 00 3e 00 47 00 45 00 35 00 2f 00
0000150 1d 00 14 00 09 00 10 00 17 00 1f 00 27 00 3d 00
0000160 55 00 65 00 5a 00 3d 00 32 00 29 00 1f 00 2c 00
0000170 1a 00 1b 00 19 00 0a 00 19 00 1b 00 25 00 29 00
0000180 38 00 4c 00 5a 00 5e 00 3c 00 30 00 28 00 25 00
0000190 22 00 1e 00 25 00 31 00 09 00 1c 00 31 00 2f 00
00001a0 26 00 31 00 35 00 4c 00 56 00 4d 00 3b 00 28 00
00001b0 29 00 29 00 26 00 27 00 49 00 10 00 19 00 29 00
00001c0 2f 00 21 00 21 00 26 00 43 00 48 00 49 00 36 00
00001d0 24 00 21 00 24 00 2a 00 30 00 45 00 0c 00 0e 00
00001e0 1f 00 1d 00 19 00 1a 00 26 00 3d 00 45 00 2d 00
00001f0 1d 00 1f 00 22 00 24 00 2d 00 30 00 35 00 1a 00
0000200 13 00 10 00 1d 00 17 00 1a 00 1d 00 28 00 29 00
0000210 1c 00 1c 00 2a 00 2c 00 2f 00 24 00 1e 00 23 00
0000220 22 00 19 00 13 00 10 00 0f 00 17 00 13 00 1a 00
0000230 1b 00 15 00 19 00 2d 00 33 00 34 00 2e 00 1b 00
0000240 21 00
0000242
```
