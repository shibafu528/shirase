package shirase

import (
	"fmt"
	"hash/fnv"
	"os"

	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func GenerateSnowflakeID() (snowflake.ID, error) {
	sf, err := getSnowflake()
	if err != nil {
		return 0, fmt.Errorf("[shirase.GenerateSnowflakeID] id generate error: %w", err)
	}
	return sf.Generate(), nil
}

func getSnowflake() (_ *snowflake.Node, err error) {
	if snowflakeNode == nil {
		var hostname uint64
		hostname, err = hashedHostname()
		if err != nil {
			return
		}
		pid := os.Getpid()
		nid := hostname ^ (uint64)(pid)

		snowflakeNode, err = snowflake.NewNode(int64(nid) & 0x3FF)
		if err != nil {
			return
		}
	}
	return snowflakeNode, nil
}

func hashedHostname() (uint64, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return 0, err
	}

	f := fnv.New64a()
	_, err = f.Write([]byte(hostname))
	if err != nil {
		return 0, err
	}

	return f.Sum64(), nil
}
