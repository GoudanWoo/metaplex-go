// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateMetadataAccountV2 is the `UpdateMetadataAccountV2` instruction.
type UpdateMetadataAccountV2 struct {
	UpdateMetadataAccountArgsV2 *UpdateMetadataAccountArgsV2

	// [0] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [1] = [SIGNER] updateAuthority
	// ··········· Update authority key
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateMetadataAccountV2InstructionBuilder creates a new `UpdateMetadataAccountV2` instruction builder.
func NewUpdateMetadataAccountV2InstructionBuilder() *UpdateMetadataAccountV2 {
	nd := &UpdateMetadataAccountV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetUpdateMetadataAccountArgsV2 sets the "updateMetadataAccountArgsV2" parameter.
func (inst *UpdateMetadataAccountV2) SetUpdateMetadataAccountArgsV2(updateMetadataAccountArgsV2 UpdateMetadataAccountArgsV2) *UpdateMetadataAccountV2 {
	inst.UpdateMetadataAccountArgsV2 = &updateMetadataAccountArgsV2
	return inst
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *UpdateMetadataAccountV2) SetMetadataAccount(metadata ag_solanago.PublicKey) *UpdateMetadataAccountV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *UpdateMetadataAccountV2) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Update authority key
func (inst *UpdateMetadataAccountV2) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *UpdateMetadataAccountV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Update authority key
func (inst *UpdateMetadataAccountV2) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst UpdateMetadataAccountV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateMetadataAccountV2),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateMetadataAccountV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateMetadataAccountV2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UpdateMetadataAccountArgsV2 == nil {
			return errors.New("UpdateMetadataAccountArgsV2 parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
	}
	return nil
}

func (inst *UpdateMetadataAccountV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateMetadataAccountV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UpdateMetadataAccountArgsV2", *inst.UpdateMetadataAccountArgsV2))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdateMetadataAccountV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UpdateMetadataAccountArgsV2` param:
	err = encoder.Encode(obj.UpdateMetadataAccountArgsV2)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateMetadataAccountV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UpdateMetadataAccountArgsV2`:
	err = decoder.Decode(&obj.UpdateMetadataAccountArgsV2)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateMetadataAccountV2Instruction declares a new UpdateMetadataAccountV2 instruction with the provided parameters and accounts.
func NewUpdateMetadataAccountV2Instruction(
	// Parameters:
	updateMetadataAccountArgsV2 UpdateMetadataAccountArgsV2,
	// Accounts:
	metadata ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey) *UpdateMetadataAccountV2 {
	return NewUpdateMetadataAccountV2InstructionBuilder().
		SetUpdateMetadataAccountArgsV2(updateMetadataAccountArgsV2).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityAccount(updateAuthority)
}
