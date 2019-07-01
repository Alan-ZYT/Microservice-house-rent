redis-server ./conf/redis.conf

fdfs_trackerd  /home/itcast/go/src/go-1/homeweb/conf/tracker.conf restart

fdfs_storaged  /home/itcast/go/src/go-1/homeweb/conf/storage.conf restart

sudo nginx
