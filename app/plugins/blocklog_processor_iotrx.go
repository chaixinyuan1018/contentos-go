package plugins

import (
	"github.com/coschain/contentos-go/app/blocklog"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/prototype"
	"github.com/jinzhu/gorm"
)

type OpProcessor func(db *gorm.DB, blockLog *blocklog.BlockLog, opIdx, trxIdx int) error
type ChangeProcessor func(db *gorm.DB, change *blocklog.StateChange, blockLog *blocklog.BlockLog, changeIdx, opIdx, trxIdx int) error



type IOTrxProcessor struct {
	tableReady bool
	opProcessors []OpProcessor
	changeProcessors []ChangeProcessor
}

func NewIOTrxProcessor() *IOTrxProcessor{
	p := &IOTrxProcessor{}
	p.addOpProcessors()
	p.addChangeProcessor()
	return p
}

func (p *IOTrxProcessor) addOpProcessors() {
	p.opProcessors = append(p.opProcessors, ProcessAccountCreateOperation, ProcessTransferOperation,
		ProcessTransferVestOperation, ProcessStakeOperation, ProcessUnStakeOperation,
		ProcessAccountUpdateOperation, ProcessVoteOperation, ProcessBpRegisterOperation, ProcessBpUpdateOperation,
		ProcessBpEnableOperation, ProcessBpVoteOperation, ProcessContractDeployOperation, ProcessContractApplyOperation,
		ProcessPostOperation, ProcessReplyOperation, ProcessConvertVestOperation, ProcessAcquireTicketOperation,
		ProcessVoteByTicketOperation)
}

func (p *IOTrxProcessor) addChangeProcessor() {
	p.changeProcessors = append(p.changeProcessors,
		ProcessContractTransferToUserChangeProcessor,
		ProcessUserToContractChangeProcessor,
		ProcessContractTransferToContractChangeProcessor)
}

func (p *IOTrxProcessor) Prepare(db *gorm.DB, blockLog *blocklog.BlockLog) (err error) {
	if !p.tableReady {
		if !db.HasTable(&iservices.IOTrxRecord{}) {
			if err = db.CreateTable(&iservices.IOTrxRecord{}).Error; err == nil {
				p.tableReady = true
			}
		} else {
			p.tableReady = true
		}
	}
	return
}

func (p *IOTrxProcessor) ProcessChange(db *gorm.DB, change *blocklog.StateChange, blockLog *blocklog.BlockLog, changeIdx, opIdx, trxIdx int) error {
	for _, changeProcessor := range p.changeProcessors {
		if err := changeProcessor(db, change, blockLog, changeIdx, opIdx, trxIdx); err != nil {
			return err
		}
	}
	return nil
}

func (p *IOTrxProcessor) ProcessOperation(db *gorm.DB, blockLog *blocklog.BlockLog, opIdx, trxIdx int) error {
	trxLog := blockLog.Transactions[trxIdx]
	if trxLog.Receipt.Status != prototype.StatusSuccess {
		return nil
	}
	for _, opProcessor := range p.opProcessors {
		if err := opProcessor(db, blockLog, opIdx, trxIdx); err != nil {
			return err
		}
	}
	return nil
}

func (p *IOTrxProcessor) Finalize(db *gorm.DB, blockLog *blocklog.BlockLog) error {
	return nil
}

func init() {
	RegisterSQLTableNamePattern("iotrx_record")
}