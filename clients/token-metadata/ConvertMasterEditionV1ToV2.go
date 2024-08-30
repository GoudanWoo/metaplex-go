// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConvertMasterEditionV1ToV2 is the `ConvertMasterEditionV1ToV2` instruction.
type ConvertMasterEditionV1ToV2 struct {

	// [0] = [WRITE] masterEdition
	// ··········· Master Record Edition V1 (pda of ['metadata', program id, master metadata mint id, 'edition'])
	//
	// [1] = [WRITE] oneTimeAuth
	// ··········· One time authorization mint
	//
	// [2] = [WRITE] printingMint
	// ··········· Printing mint
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConvertMasterEditionV1ToV2InstructionBuilder creates a new `ConvertMasterEditionV1ToV2` instruction builder.
func NewConvertMasterEditionV1ToV2InstructionBuilder() *ConvertMasterEditionV1ToV2 {
	nd := &ConvertMasterEditionV1ToV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMasterEditionAccount sets the "masterEdition" account.
// Master Record Edition V1 (pda of ['metadata', program id, master metadata mint id, 'edition'])
func (inst *ConvertMasterEditionV1ToV2) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *ConvertMasterEditionV1ToV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
// Master Record Edition V1 (pda of ['metadata', program id, master metadata mint id, 'edition'])
func (inst *ConvertMasterEditionV1ToV2) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOneTimeAuthAccount sets the "oneTimeAuth" account.
// One time authorization mint
func (inst *ConvertMasterEditionV1ToV2) SetOneTimeAuthAccount(oneTimeAuth ag_solanago.PublicKey) *ConvertMasterEditionV1ToV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(oneTimeAuth).WRITE()
	return inst
}

// GetOneTimeAuthAccount gets the "oneTimeAuth" account.
// One time authorization mint
func (inst *ConvertMasterEditionV1ToV2) GetOneTimeAuthAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPrintingMintAccount sets the "printingMint" account.
// Printing mint
func (inst *ConvertMasterEditionV1ToV2) SetPrintingMintAccount(printingMint ag_solanago.PublicKey) *ConvertMasterEditionV1ToV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(printingMint).WRITE()
	return inst
}

// GetPrintingMintAccount gets the "printingMint" account.
// Printing mint
func (inst *ConvertMasterEditionV1ToV2) GetPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst ConvertMasterEditionV1ToV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ConvertMasterEditionV1ToV2),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConvertMasterEditionV1ToV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConvertMasterEditionV1ToV2) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OneTimeAuth is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PrintingMint is not set")
		}
	}
	return nil
}

func (inst *ConvertMasterEditionV1ToV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConvertMasterEditionV1ToV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("masterEdition", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  oneTimeAuth", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" printingMint", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj ConvertMasterEditionV1ToV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ConvertMasterEditionV1ToV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewConvertMasterEditionV1ToV2Instruction declares a new ConvertMasterEditionV1ToV2 instruction with the provided parameters and accounts.
func NewConvertMasterEditionV1ToV2Instruction(
	// Accounts:
	masterEdition ag_solanago.PublicKey,
	oneTimeAuth ag_solanago.PublicKey,
	printingMint ag_solanago.PublicKey) *ConvertMasterEditionV1ToV2 {
	return NewConvertMasterEditionV1ToV2InstructionBuilder().
		SetMasterEditionAccount(masterEdition).
		SetOneTimeAuthAccount(oneTimeAuth).
		SetPrintingMintAccount(printingMint)
}
