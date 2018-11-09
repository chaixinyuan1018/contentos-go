package types

import (
	comm "github.com/coschain/contentos-go/p2p/depend/common"
	"github.com/coschain/contentos-go/p2p/depend/core/types"
	"github.com/coschain/contentos-go/p2p/common"
)

// Transaction message
type Trn struct {
	Txn *types.Transaction
}

//Serialize message payload
func (this Trn) Serialization(sink *comm.ZeroCopySink) error {
	return this.Txn.Serialization(sink)
}

func (this *Trn) CmdType() string {
	return common.TX_TYPE
}

//Deserialize message payload
func (this *Trn) Deserialization(source *comm.ZeroCopySource) error {
	tx := &types.Transaction{}
	err := tx.Deserialization(source)
	if err != nil {
		return err
	}

	this.Txn = tx
	return nil
}
