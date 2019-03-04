package horizonclient

import (
	"context"
	"net/http"

	"github.com/stellar/go/support/app"
	"github.com/stellar/go/support/errors"
)

func sendRequest(hr HorizonRequest, c Client, a interface{}) (err error) {
	endpoint, err := hr.BuildUrl()
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", c.HorizonURL+endpoint, nil)
	if err != nil {
		return errors.Wrap(err, "Error creating HTTP request")
	}
	req.Header.Set("X-Client-Name", "go-stellar-sdk")
	// to do: Confirm if there is a different way to set version. Not sure about this, since we dont build the sdk into an executable file.
	// Do we currently track sdk versions differently?
	req.Header.Set("X-Client-Version", app.Version())

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &a)
	return
}

// AccountDetail returns information for a single account.
// See https://www.stellar.org/developers/horizon/reference/endpoints/accounts-single.html
func (c *Client) AccountDetail(request AccountRequest) (account Account, err error) {
	if request.AccountId == "" {
		err = errors.New("No account ID provided")
	}

	if err != nil {
		return
	}

	err = sendRequest(request, *c, &account)
	return
}

// AccountData returns a single data associated with a given account
// See https://www.stellar.org/developers/horizon/reference/endpoints/data-for-account.html
func (c *Client) AccountData(request AccountRequest) (accountData AccountData, err error) {
	if request.AccountId == "" || request.DataKey == "" {
		err = errors.New("Too few parameters")
	}

	if err != nil {
		return
	}

	err = sendRequest(request, *c, &accountData)
	return
}

// Effects returns effects(https://www.stellar.org/developers/horizon/reference/resources/effect.html)
// It can be used to return effects for an account, a ledger, an operation, a transaction and all effects on the network.
func (c *Client) Effects(request EffectRequest) (effects EffectsPage, err error) {
	err = sendRequest(request, *c, &effects)
	return
}

// Assets returns asset information.
// See https://www.stellar.org/developers/horizon/reference/endpoints/assets-all.html
func (c *Client) Assets(request AssetRequest) (assets AssetsPage, err error) {
	err = sendRequest(request, *c, &assets)
	return
}

func (c *Client) Stream(request StreamRequest, ctx context.Context, handler func(interface{})) (err error) {

	err = request.Stream(c.HorizonURL, ctx, handler)
	return
}
