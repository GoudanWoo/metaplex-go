// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ApproveCollectionAuthority is the `ApproveCollectionAuthority` instruction.
type ApproveCollectionAuthority struct {

	// [0] = [WRITE] collectionAuthorityRecord
	// ··········· Collection Authority Record PDA
	//
	// [1] = [] newCollectionAuthority
	// ··········· A Collection Authority
	//
	// [2] = [WRITE, SIGNER] updateAuthority
	// ··········· Update Authority of Collection NFT
	//
	// [3] = [WRITE, SIGNER] payer
	// ··········· Payer
	//
	// [4] = [] metadata
	// ··········· Collection Metadata account
	//
	// [5] = [] mint
	// ··········· Mint of Collection Metadata
	//
	// [6] = [] systemProgram
	// ··········· System program
	//
	// [7] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewApproveCollectionAuthorityInstructionBuilder creates a new `ApproveCollectionAuthority` instruction builder.
func NewApproveCollectionAuthorityInstructionBuilder() *ApproveCollectionAuthority {
	nd := &ApproveCollectionAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetCollectionAuthorityRecordAccount sets the "collectionAuthorityRecord" account.
// Collection Authority Record PDA
func (inst *ApproveCollectionAuthority) SetCollectionAuthorityRecordAccount(collectionAuthorityRecord ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(collectionAuthorityRecord).WRITE()
	return inst
}

// GetCollectionAuthorityRecordAccount gets the "collectionAuthorityRecord" account.
// Collection Authority Record PDA
func (inst *ApproveCollectionAuthority) GetCollectionAuthorityRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetNewCollectionAuthorityAccount sets the "newCollectionAuthority" account.
// A Collection Authority
func (inst *ApproveCollectionAuthority) SetNewCollectionAuthorityAccount(newCollectionAuthority ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(newCollectionAuthority)
	return inst
}

// GetNewCollectionAuthorityAccount gets the "newCollectionAuthority" account.
// A Collection Authority
func (inst *ApproveCollectionAuthority) GetNewCollectionAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Update Authority of Collection NFT
func (inst *ApproveCollectionAuthority) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).WRITE().SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Update Authority of Collection NFT
func (inst *ApproveCollectionAuthority) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *ApproveCollectionAuthority) SetPayerAccount(payer ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *ApproveCollectionAuthority) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMetadataAccount sets the "metadata" account.
// Collection Metadata account
func (inst *ApproveCollectionAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Collection Metadata account
func (inst *ApproveCollectionAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMintAccount sets the "mint" account.
// Mint of Collection Metadata
func (inst *ApproveCollectionAuthority) SetMintAccount(mint ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// Mint of Collection Metadata
func (inst *ApproveCollectionAuthority) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *ApproveCollectionAuthority) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *ApproveCollectionAuthority) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *ApproveCollectionAuthority) SetRentAccount(rent ag_solanago.PublicKey) *ApproveCollectionAuthority {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account (optional).
// Rent info
func (inst *ApproveCollectionAuthority) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst ApproveCollectionAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ApproveCollectionAuthority),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ApproveCollectionAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ApproveCollectionAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CollectionAuthorityRecord is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.NewCollectionAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}

		// [7] = Rent is optional

	}
	return nil
}

func (inst *ApproveCollectionAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ApproveCollectionAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("collectionAuthorityRecord", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   newCollectionAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          updateAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                     mint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            systemProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                     rent", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj ApproveCollectionAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ApproveCollectionAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewApproveCollectionAuthorityInstruction declares a new ApproveCollectionAuthority instruction with the provided parameters and accounts.
func NewApproveCollectionAuthorityInstruction(
	// Accounts:
	collectionAuthorityRecord ag_solanago.PublicKey,
	newCollectionAuthority ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *ApproveCollectionAuthority {
	return NewApproveCollectionAuthorityInstructionBuilder().
		SetCollectionAuthorityRecordAccount(collectionAuthorityRecord).
		SetNewCollectionAuthorityAccount(newCollectionAuthority).
		SetUpdateAuthorityAccount(updateAuthority).
		SetPayerAccount(payer).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
