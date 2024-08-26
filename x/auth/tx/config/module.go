package tx

import (
	"context"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/x/auth/ante"
)

// TxValidator is appmodulev2.HasTxValidator without the AppModule requirement.
type TxValidator = func(context.Context, transaction.Tx) error

var (
	_ appmodulev2.AppModule                      = AppModule{}
	_ appmodulev2.HasTxValidator[transaction.Tx] = AppModule{}
)

// AppModule is a module that only implements tx validators.
// The goal of this module is to allow extensible registration of tx validators provided by chains without requiring a new modules.
// Additionally, it registers tx validators that do not really have a place in other modules.
// This module is only useful for chains using server/v2. Ante/Post handlers are setup via baseapp options in depinject.
type AppModule struct {
	sigVerification ante.SigVerificationDecorator
	// txValidators contains tx validator that can be injected into the module via depinject.
	// tx validators should be module based, but it can happen that you do not want to create a new module
	// and simply depinject-it.
	txValidators []TxValidator
}

// NewAppModule creates a new AppModule object.
func NewAppModule(
	sigVerification ante.SigVerificationDecorator,
	txValidators ...TxValidator,
) AppModule {
	return AppModule{
		sigVerification: sigVerification,
		txValidators:    txValidators,
	}
}

// IsAppModule implements appmodule.AppModule.
func (a AppModule) IsAppModule() {}

// IsOnePerModuleType implements appmodule.AppModule.
func (a AppModule) IsOnePerModuleType() {}

// TxValidator implements appmodule.HasTxValidator.
func (a AppModule) TxValidator(ctx context.Context, tx transaction.Tx) error {
	for _, validator := range a.txValidators {
		if err := validator(ctx, tx); err != nil {
			return err
		}
	}

	return a.sigVerification.ValidateTx(ctx, tx)
}
