// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_ccip_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateRouter is the `updateRouter` instruction.
type UpdateRouter struct {
	NewRouter *ag_solanago.PublicKey

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateRouterInstructionBuilder creates a new `UpdateRouter` instruction builder.
func NewUpdateRouterInstructionBuilder() *UpdateRouter {
	nd := &UpdateRouter{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewRouter sets the "newRouter" parameter.
func (inst *UpdateRouter) SetNewRouter(newRouter ag_solanago.PublicKey) *UpdateRouter {
	inst.NewRouter = &newRouter
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *UpdateRouter) SetStateAccount(state ag_solanago.PublicKey) *UpdateRouter {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateRouter) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateRouter) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateRouter {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateRouter) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst UpdateRouter) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateRouter,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateRouter) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateRouter) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewRouter == nil {
			return errors.New("NewRouter parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdateRouter) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateRouter")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewRouter", *inst.NewRouter))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj UpdateRouter) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewRouter` param:
	err = encoder.Encode(obj.NewRouter)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateRouter) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewRouter`:
	err = decoder.Decode(&obj.NewRouter)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateRouterInstruction declares a new UpdateRouter instruction with the provided parameters and accounts.
func NewUpdateRouterInstruction(
	// Parameters:
	newRouter ag_solanago.PublicKey,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateRouter {
	return NewUpdateRouterInstructionBuilder().
		SetNewRouter(newRouter).
		SetStateAccount(state).
		SetAuthorityAccount(authority)
}
