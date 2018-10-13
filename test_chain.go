package chain_reorganization

import (
	"github.com/copernet/copernicus/model/block"
	"github.com/copernet/copernicus/model/blockindex"
)

func createBlock() {
	header := block.NewBlockHeader()
	index := blockindex.NewBlockIndex(header)
	index.Height = int32(height)
}
