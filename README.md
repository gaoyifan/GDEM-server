# GDEM-server
The height data server for GDEM

# Install via docker

	docker run --restart=always --name=redis -v /path/to/redis:/data -d redis:3
    docker run --restart=always --name=gdemd -v /path/to/map:/srv/map --link redis:redis -d -p 8000:8000 gaoyfian/gdem-server

# Download Map Data

You can downlaod the GDEM v2 Map Data(global data) which is converted to png format:

    http://pan.baidu.com/s/1pJBT1qz

Or you can just downlaod the data you need.

# API Format

	http://hostname/i/j/zoom/size/
	
* i: abscissa
* j: ordinate
* zoom: Zoom levels ([Reference](http://wiki.openstreetmap.org/wiki/Slippy_map_tilenames))
* size: The resolution of the map block. 2^size * 2^size

