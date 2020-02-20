package doubledpool

import (
	"go.intra.xiaojukeji.com/gulfstream/dConnPool/redigo"
	"sync"
	dpool "go.intra.xiaojukeji.com/gulfstream/dConnPool"

	"time"
	"helper/nodemgr_ex"
	"runtime/debug"
	logger "github.com/shengkehua/xlog4go"
)

const (
	BUFFER_SIZE = 2
)

type DoubleBufferDPool struct {
	pools   [BUFFER_SIZE]*redigo.DPool
	lock    [BUFFER_SIZE]sync.RWMutex
	serverName string
	hosts   []string
	current int
}

func (buffer *DoubleBufferDPool) Pool() *redigo.DPool {
	index := buffer.current
	buffer.lock[index].RLock()
	defer buffer.lock[index].RUnlock()

	return buffer.pools[index]
}

func (buffer *DoubleBufferDPool) Close() {
	for _, pool := range buffer.pools {
		pool.Close()
	}
}

func (buffer *DoubleBufferDPool) NewDPools(servernName string, connFactory redigo.PooledConnFactory, poolConfig dpool.PoolConfig) (err error) {
	servers, err := nodemgr_ex.GetAllNodesHosts(servernName)
	if err != nil {
		return err
	}

	for i := range buffer.pools {
		buffer.pools[i] = redigo.NewDPool(servers, connFactory, poolConfig)
	}

	buffer.serverName = servernName
	buffer.hosts = servers[:]
	buffer.current = 0
	return
}

func (buffer *DoubleBufferDPool) HasDiff(oldHosts, newHosts []string) bool {
	if newHosts == nil {
		return false
	}

	if len(oldHosts) != len(newHosts) {
		return true
	}

	oldMap := make(map[string]int8, len(oldHosts))
	for _, host := range oldHosts {
		oldMap[host] += 1
	}
	for _, host := range newHosts {
		if _, ok := oldMap[host]; !ok {
			return true
		} else {
			oldMap[host] -= 1
		}
	}
	for _, count := range oldMap {
		if count != 0 {
			return true
		}
	}

	return false
}

func (buffer *DoubleBufferDPool) ReFresh(poolConfig dpool.PoolConfig, factory redigo.PooledConnFactory, interval int) {
	// set recover
	defer func() {
		if err := recover(); err != nil {
			logger.Error("abort, unknown error, reason:%v, stack:%s", err, string(debug.Stack()))
		}
	}()

	var timer *time.Ticker = time.NewTicker(time.Duration(interval) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			servers, err := nodemgr_ex.GetAllNodesHosts(buffer.serverName)
			if err != nil || servers == nil || len(servers) == 0 {
				logger.Warn("DoubleBufferDPool ReFresh GetAllNodesHosts failed, err:%v", err)
				continue
			}

			if buffer.hosts != nil && !buffer.HasDiff(buffer.hosts, servers) {
				continue
			}
			buffer.hosts = servers[:]

			// get servers and new
			index := (buffer.current + 1) % BUFFER_SIZE

			buffer.lock[index].Lock()
			buffer.pools[index].Close()
			buffer.pools[index] = redigo.NewDPool(servers, factory, poolConfig)
			buffer.lock[index].Unlock()
			buffer.current = index
		}
	}


}
