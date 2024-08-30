// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Burn is the `Burn` instruction.
type Burn struct {
	BurnArgs BurnArgs

	// [0] = [WRITE, SIGNER] authority
	// ··········· Asset owner or Utility delegate
	//
	// [1] = [WRITE] collectionMetadata
	// ··········· Metadata of the Collection
	//
	// [2] = [WRITE] metadata
	// ··········· Metadata (pda of ['metadata', program id, mint id])
	//
	// [3] = [WRITE] edition
	// ··········· Edition of the asset
	//
	// [4] = [WRITE] mint
	// ··········· Mint of token asset
	//
	// [5] = [WRITE] token
	// ··········· Token account to close
	//
	// [6] = [WRITE] masterEdition
	// ··········· Master edition account
	//
	// [7] = [] masterEditionMint
	// ··········· Master edition mint of the asset
	//
	// [8] = [] masterEditionToken
	// ··········· Master edition token account
	//
	// [9] = [WRITE] editionMarker
	// ··········· Edition marker account
	//
	// [10] = [WRITE] tokenRecord
	// ··········· Token record account
	//
	// [11] = [] systemProgram
	// ··········· System program
	//
	// [12] = [] sysvarInstructions
	// ··········· Instructions sysvar account
	//
	// [13] = [] splTokenProgram
	// ··········· SPL Token Program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBurnInstructionBuilder creates a new `Burn` instruction builder.
func NewBurnInstructionBuilder() *Burn {
	nd := &Burn{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetBurnArgs sets the "burnArgs" parameter.
func (inst *Burn) SetBurnArgs(burnArgs BurnArgs) *Burn {
	inst.BurnArgs = burnArgs
	return inst
}

// SetAuthorityAccount sets the "authority" account.
// Asset owner or Utility delegate
func (inst *Burn) SetAuthorityAccount(authority ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Asset owner or Utility delegate
func (inst *Burn) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
// Metadata of the Collection
func (inst *Burn) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collectionMetadata).WRITE()
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account (optional).
// Metadata of the Collection
func (inst *Burn) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata (pda of ['metadata', program id, mint id])
func (inst *Burn) SetMetadataAccount(metadata ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata (pda of ['metadata', program id, mint id])
func (inst *Burn) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEditionAccount sets the "edition" account.
// Edition of the asset
func (inst *Burn) SetEditionAccount(edition ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(edition).WRITE()
	return inst
}

// GetEditionAccount gets the "edition" account (optional).
// Edition of the asset
func (inst *Burn) GetEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintAccount sets the "mint" account.
// Mint of token asset
func (inst *Burn) SetMintAccount(mint ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
// Mint of token asset
func (inst *Burn) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenAccount sets the "token" account.
// Token account to close
func (inst *Burn) SetTokenAccount(token ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(token).WRITE()
	return inst
}

// GetTokenAccount gets the "token" account.
// Token account to close
func (inst *Burn) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMasterEditionAccount sets the "masterEdition" account.
// Master edition account
func (inst *Burn) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account (optional).
// Master edition account
func (inst *Burn) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMasterEditionMintAccount sets the "masterEditionMint" account.
// Master edition mint of the asset
func (inst *Burn) SetMasterEditionMintAccount(masterEditionMint ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(masterEditionMint)
	return inst
}

// GetMasterEditionMintAccount gets the "masterEditionMint" account (optional).
// Master edition mint of the asset
func (inst *Burn) GetMasterEditionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMasterEditionTokenAccount sets the "masterEditionToken" account.
// Master edition token account
func (inst *Burn) SetMasterEditionTokenAccount(masterEditionToken ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(masterEditionToken)
	return inst
}

// GetMasterEditionTokenAccount gets the "masterEditionToken" account (optional).
// Master edition token account
func (inst *Burn) GetMasterEditionTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetEditionMarkerAccount sets the "editionMarker" account.
// Edition marker account
func (inst *Burn) SetEditionMarkerAccount(editionMarker ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(editionMarker).WRITE()
	return inst
}

// GetEditionMarkerAccount gets the "editionMarker" account (optional).
// Edition marker account
func (inst *Burn) GetEditionMarkerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenRecordAccount sets the "tokenRecord" account.
// Token record account
func (inst *Burn) SetTokenRecordAccount(tokenRecord ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenRecord).WRITE()
	return inst
}

// GetTokenRecordAccount gets the "tokenRecord" account (optional).
// Token record account
func (inst *Burn) GetTokenRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *Burn) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *Burn) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *Burn) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *Burn) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// SPL Token Program
func (inst *Burn) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account.
// SPL Token Program
func (inst *Burn) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst Burn) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_Burn),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Burn) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Burn) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.BurnArgs == nil {
			return errors.New("BurnArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}

		// [1] = CollectionMetadata is optional

		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Metadata is not set")
		}

		// [3] = Edition is optional

		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Token is not set")
		}

		// [6] = MasterEdition is optional

		// [7] = MasterEditionMint is optional

		// [8] = MasterEditionToken is optional

		// [9] = EditionMarker is optional

		// [10] = TokenRecord is optional

		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SplTokenProgram is not set")
		}
	}
	return nil
}

func (inst *Burn) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Burn")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("BurnArgs", inst.BurnArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("collectionMetadata", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          metadata", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           edition", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("              mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             token", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     masterEdition", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta(" masterEditionMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("masterEditionToken", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("     editionMarker", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       tokenRecord", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("sysvarInstructions", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("   splTokenProgram", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj Burn) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BurnArgs` param:
	{
		tmp := burnArgsContainer{}
		switch realvalue := obj.BurnArgs.(type) {
		case *BurnArgsV1:
			tmp.Enum = 0
			tmp.V1 = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	return nil
}
func (obj *Burn) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BurnArgs`:
	{
		tmp := new(burnArgsContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.BurnArgs = &tmp.V1
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	return nil
}

// NewBurnInstruction declares a new Burn instruction with the provided parameters and accounts.
func NewBurnInstruction(
	// Parameters:
	burnArgs BurnArgs,
	// Accounts:
	authority ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	edition ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	masterEditionMint ag_solanago.PublicKey,
	masterEditionToken ag_solanago.PublicKey,
	editionMarker ag_solanago.PublicKey,
	tokenRecord ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey) *Burn {
	return NewBurnInstructionBuilder().
		SetBurnArgs(burnArgs).
		SetAuthorityAccount(authority).
		SetCollectionMetadataAccount(collectionMetadata).
		SetMetadataAccount(metadata).
		SetEditionAccount(edition).
		SetMintAccount(mint).
		SetTokenAccount(token).
		SetMasterEditionAccount(masterEdition).
		SetMasterEditionMintAccount(masterEditionMint).
		SetMasterEditionTokenAccount(masterEditionToken).
		SetEditionMarkerAccount(editionMarker).
		SetTokenRecordAccount(tokenRecord).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetSplTokenProgramAccount(splTokenProgram)
}
