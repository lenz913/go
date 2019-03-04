package txnbuild

import (
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

// SetOptions represents the Stellar set options operation. See
// https://www.stellar.org/developers/guides/concepts/list-of-operations.html
type SetOptions struct {
	destAccountID        xdr.AccountId
	InflationDestination string
	xdrOp                xdr.SetOptionsOp
}

// BuildXDR for SetOptions returns a fully configured XDR Operation.
func (so *SetOptions) BuildXDR() (xdr.Operation, error) {
	// Inflation Destination. Once set, there's no way to unset.
	if so.InflationDestination != "" {
		err := so.destAccountID.SetAddress(so.InflationDestination)
		if err != nil {
			return xdr.Operation{}, errors.Wrap(err, "Failed to set inflation destination address")
		}
		so.xdrOp.InflationDest = &so.destAccountID
	}

	opType := xdr.OperationTypeSetOptions
	body, err := xdr.NewOperationBody(opType, so.xdrOp)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to build XDR OperationBody")
	}

	return xdr.Operation{Body: body}, nil
}
