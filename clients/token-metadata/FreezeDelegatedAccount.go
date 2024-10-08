// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FreezeDelegatedAccount is the `FreezeDelegatedAccount` instruction.
type FreezeDelegatedAccount struct {

	// [0] = [WRITE, SIGNER] delegate
	// ··········· Delegate
	//
	// [1] = [WRITE] tokenAccount
	// ··········· Token account to freeze
	//
	// [2] = [] edition
	// ··········· Edition
	//
	// [3] = [] mint
	// ··········· Token mint
	//
	// [4] = [] tokenProgram
	// ··········· Token Program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFreezeDelegatedAccountInstructionBuilder creates a new `FreezeDelegatedAccount` instruction builder.
func NewFreezeDelegatedAccountInstructionBuilder() *FreezeDelegatedAccount {
	nd := &FreezeDelegatedAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetDelegateAccount sets the "delegate" account.
// Delegate
func (inst *FreezeDelegatedAccount) SetDelegateAccount(delegate ag_solanago.PublicKey) *FreezeDelegatedAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(delegate).WRITE().SIGNER()
	return inst
}

// GetDelegateAccount gets the "delegate" account.
// Delegate
func (inst *FreezeDelegatedAccount) GetDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenAccountAccount sets the "tokenAccount" account.
// Token account to freeze
func (inst *FreezeDelegatedAccount) SetTokenAccountAccount(tokenAccount ag_solanago.PublicKey) *FreezeDelegatedAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenAccount).WRITE()
	return inst
}

// GetTokenAccountAccount gets the "tokenAccount" account.
// Token account to freeze
func (inst *FreezeDelegatedAccount) GetTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEditionAccount sets the "edition" account.
// Edition
func (inst *FreezeDelegatedAccount) SetEditionAccount(edition ag_solanago.PublicKey) *FreezeDelegatedAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(edition)
	return inst
}

// GetEditionAccount gets the "edition" account.
// Edition
func (inst *FreezeDelegatedAccount) GetEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
// Token mint
func (inst *FreezeDelegatedAccount) SetMintAccount(mint ag_solanago.PublicKey) *FreezeDelegatedAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// Token mint
func (inst *FreezeDelegatedAccount) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token Program
func (inst *FreezeDelegatedAccount) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *FreezeDelegatedAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token Program
func (inst *FreezeDelegatedAccount) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst FreezeDelegatedAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_FreezeDelegatedAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FreezeDelegatedAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FreezeDelegatedAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Delegate is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Edition is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *FreezeDelegatedAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FreezeDelegatedAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    delegate", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       token", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     edition", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("tokenProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj FreezeDelegatedAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *FreezeDelegatedAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewFreezeDelegatedAccountInstruction declares a new FreezeDelegatedAccount instruction with the provided parameters and accounts.
func NewFreezeDelegatedAccountInstruction(
	// Accounts:
	delegate ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey,
	edition ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *FreezeDelegatedAccount {
	return NewFreezeDelegatedAccountInstructionBuilder().
		SetDelegateAccount(delegate).
		SetTokenAccountAccount(tokenAccount).
		SetEditionAccount(edition).
		SetMintAccount(mint).
		SetTokenProgramAccount(tokenProgram)
}
