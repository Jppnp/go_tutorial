package models

import "go/tutorial/models/entity"

type AccountingStageModel struct {
	OldAccouting  entity.Accounting
	NewAccounting entity.Accounting
}
