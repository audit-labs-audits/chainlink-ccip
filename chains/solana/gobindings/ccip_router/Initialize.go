// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialization Flow //
// Initializes the CCIP Router.
//
// The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for initialization.
// * `svm_chain_selector` - The chain selector for SVM.
// * `fee_aggregator` - The public key of the fee aggregator.
// * `fee_quoter` - The public key of the fee quoter.
// * `link_token_mint` - The public key of the LINK token mint.
// * `rmn_remote` - The public key of the RMN remote.
type Initialize struct {
	SvmChainSelector *uint64
	FeeAggregator    *ag_solanago.PublicKey
	FeeQuoter        *ag_solanago.PublicKey
	LinkTokenMint    *ag_solanago.PublicKey
	RmnRemote        *ag_solanago.PublicKey

	// [0] = [WRITE] config
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [] systemProgram
	//
	// [3] = [] program
	//
	// [4] = [] programData
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetSvmChainSelector sets the "svmChainSelector" parameter.
func (inst *Initialize) SetSvmChainSelector(svmChainSelector uint64) *Initialize {
	inst.SvmChainSelector = &svmChainSelector
	return inst
}

// SetFeeAggregator sets the "feeAggregator" parameter.
func (inst *Initialize) SetFeeAggregator(feeAggregator ag_solanago.PublicKey) *Initialize {
	inst.FeeAggregator = &feeAggregator
	return inst
}

// SetFeeQuoter sets the "feeQuoter" parameter.
func (inst *Initialize) SetFeeQuoter(feeQuoter ag_solanago.PublicKey) *Initialize {
	inst.FeeQuoter = &feeQuoter
	return inst
}

// SetLinkTokenMint sets the "linkTokenMint" parameter.
func (inst *Initialize) SetLinkTokenMint(linkTokenMint ag_solanago.PublicKey) *Initialize {
	inst.LinkTokenMint = &linkTokenMint
	return inst
}

// SetRmnRemote sets the "rmnRemote" parameter.
func (inst *Initialize) SetRmnRemote(rmnRemote ag_solanago.PublicKey) *Initialize {
	inst.RmnRemote = &rmnRemote
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Initialize) SetConfigAccount(config ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Initialize) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetProgramAccount sets the "program" account.
func (inst *Initialize) SetProgramAccount(program ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Initialize) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetProgramDataAccount sets the "programData" account.
func (inst *Initialize) SetProgramDataAccount(programData ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "programData" account.
func (inst *Initialize) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SvmChainSelector == nil {
			return errors.New("SvmChainSelector parameter is not set")
		}
		if inst.FeeAggregator == nil {
			return errors.New("FeeAggregator parameter is not set")
		}
		if inst.FeeQuoter == nil {
			return errors.New("FeeQuoter parameter is not set")
		}
		if inst.LinkTokenMint == nil {
			return errors.New("LinkTokenMint parameter is not set")
		}
		if inst.RmnRemote == nil {
			return errors.New("RmnRemote parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Program is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ProgramData is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("SvmChainSelector", *inst.SvmChainSelector))
						paramsBranch.Child(ag_format.Param("   FeeAggregator", *inst.FeeAggregator))
						paramsBranch.Child(ag_format.Param("       FeeQuoter", *inst.FeeQuoter))
						paramsBranch.Child(ag_format.Param("   LinkTokenMint", *inst.LinkTokenMint))
						paramsBranch.Child(ag_format.Param("       RmnRemote", *inst.RmnRemote))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("      program", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("  programData", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SvmChainSelector` param:
	err = encoder.Encode(obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Serialize `FeeAggregator` param:
	err = encoder.Encode(obj.FeeAggregator)
	if err != nil {
		return err
	}
	// Serialize `FeeQuoter` param:
	err = encoder.Encode(obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Serialize `LinkTokenMint` param:
	err = encoder.Encode(obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Serialize `RmnRemote` param:
	err = encoder.Encode(obj.RmnRemote)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SvmChainSelector`:
	err = decoder.Decode(&obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `FeeAggregator`:
	err = decoder.Decode(&obj.FeeAggregator)
	if err != nil {
		return err
	}
	// Deserialize `FeeQuoter`:
	err = decoder.Decode(&obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Deserialize `LinkTokenMint`:
	err = decoder.Decode(&obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `RmnRemote`:
	err = decoder.Decode(&obj.RmnRemote)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	svmChainSelector uint64,
	feeAggregator ag_solanago.PublicKey,
	feeQuoter ag_solanago.PublicKey,
	linkTokenMint ag_solanago.PublicKey,
	rmnRemote ag_solanago.PublicKey,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	program ag_solanago.PublicKey,
	programData ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetSvmChainSelector(svmChainSelector).
		SetFeeAggregator(feeAggregator).
		SetFeeQuoter(feeQuoter).
		SetLinkTokenMint(linkTokenMint).
		SetRmnRemote(rmnRemote).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData)
}
