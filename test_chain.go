package chain_reorganization

import (
	"github.com/copernet/copernicus/model/block"
	"github.com/copernet/copernicus/model/blockindex"
	"github.com/copernet/copernicus/model/tx"
	"github.com/copernet/copernicus/model/utxo"
	"github.com/copernet/copernicus/model/outpoint"
	"github.com/copernet/copernicus/util"
	"github.com/copernet/copernicus/model/undo"
	"github.com/copernet/copernicus/model"
	"github.com/copernet/copernicus/logic/lundo"
	"github.com/copernet/copernicus/persist/db"
)

func reorgBlock(blocks *block.Block, coinMap *utxo.CoinsMap, undos *undo.BlockUndo, params *model.BitcoinParams, height int) {
	header := block.NewBlockHeader()
	index := blockindex.NewBlockIndex(header)
	index.Height = int32(height)
	lundo.ApplyBlockUndo(undos, blocks, coinMap)
}

func AddCoins(txs *tx.Tx, coinMap *utxo.CoinsMap, height int32) {
	isCoinbase := txs.IsCoinBase()
	txid := txs.GetHash()
	for idx, out := range txs.GetOuts() {
		op := outpoint.NewOutPoint(txid, uint32(idx))
		coin := utxo.NewCoin(out, height, isCoinbase)
		coinMap.AddCoin(op, coin, false)
	}
}

func accessCoin(coinMap *utxo.CoinsMap, txid util.Hash) bool {
	return !coinMap.AccessCoin(outpoint.NewOutPoint(txid, 0)).IsSpent()
}

func initCoin()  {
	config := utxo.UtxoConfig{Do: &db.DBOption{FilePath: "/tmp/undotest", CacheSize: 10000}}
	utxo.InitUtxoLruTip(&config)
}