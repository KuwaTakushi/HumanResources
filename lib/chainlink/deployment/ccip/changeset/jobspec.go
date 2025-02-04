package changeset

import (
	"github.com/pkg/errors"
	"github.com/smartcontractkit/ccip-owner-contracts/pkg/proposal/timelock"

	"github.com/smartcontractkit/chainlink/deployment"
)

func Jobspec(env deployment.Environment, _ any) (deployment.ChangesetOutput, error) {
	js, err := NewCCIPJobSpecs(env.NodeIDs, env.Offchain)
	if err != nil {
		return deployment.ChangesetOutput{}, errors.Wrapf(err, "failed to create job specs")
	}
	return deployment.ChangesetOutput{
		Proposals:   []timelock.MCMSWithTimelockProposal{},
		AddressBook: deployment.NewMemoryAddressBook(),
		JobSpecs:    js,
	}, nil
}
