// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MintNewEditionFromMasterEditionViaVaultProxy is the `MintNewEditionFromMasterEditionViaVaultProxy` instruction.
type MintNewEditionFromMasterEditionViaVaultProxy struct {
	MintNewEditionFromMasterEditionViaTokenArgs *MintNewEditionFromMasterEditionViaTokenArgs

	// [0] = [WRITE] newMetadata
	// ··········· New Metadata key (pda of ['metadata', program id, mint id])
	//
	// [1] = [WRITE] newEdition
	// ··········· New Edition (pda of ['metadata', program id, mint id, 'edition'])
	//
	// [2] = [WRITE] masterEdition
	// ··········· Master Record Edition V2 (pda of ['metadata', program id, master metadata mint id, 'edition']
	//
	// [3] = [WRITE] newMint
	// ··········· Mint of new token - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
	//
	// [4] = [WRITE] editionMarkPda
	// ··········· Edition pda to mark creation - will be checked for pre-existence. (pda of ['metadata', program id, master metadata mint id, 'edition', edition_number]) where edition_number is NOT the edition number you pass in args but actually edition_number = floor(edition/EDITION_MARKER_BIT_SIZE).
	//
	// [5] = [SIGNER] newMintAuthority
	// ··········· Mint authority of new mint
	//
	// [6] = [WRITE, SIGNER] payer
	// ··········· payer
	//
	// [7] = [SIGNER] vaultAuthority
	// ··········· Vault authority
	//
	// [8] = [] safetyDepositStore
	// ··········· Safety deposit token store account
	//
	// [9] = [] safetyDepositBox
	// ··········· Safety deposit box
	//
	// [10] = [] vault
	// ··········· Vault
	//
	// [11] = [] newMetadataUpdateAuthority
	// ··········· Update authority info for new metadata
	//
	// [12] = [] metadata
	// ··········· Master record metadata account
	//
	// [13] = [] tokenProgram
	// ··········· Token program
	//
	// [14] = [] tokenVaultProgram
	// ··········· Token vault program
	//
	// [15] = [] systemProgram
	// ··········· System program
	//
	// [16] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintNewEditionFromMasterEditionViaVaultProxyInstructionBuilder creates a new `MintNewEditionFromMasterEditionViaVaultProxy` instruction builder.
func NewMintNewEditionFromMasterEditionViaVaultProxyInstructionBuilder() *MintNewEditionFromMasterEditionViaVaultProxy {
	nd := &MintNewEditionFromMasterEditionViaVaultProxy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 17),
	}
	return nd
}

// SetMintNewEditionFromMasterEditionViaTokenArgs sets the "mintNewEditionFromMasterEditionViaTokenArgs" parameter.
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetMintNewEditionFromMasterEditionViaTokenArgs(mintNewEditionFromMasterEditionViaTokenArgs MintNewEditionFromMasterEditionViaTokenArgs) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.MintNewEditionFromMasterEditionViaTokenArgs = &mintNewEditionFromMasterEditionViaTokenArgs
	return inst
}

// SetNewMetadataAccount sets the "newMetadata" account.
// New Metadata key (pda of ['metadata', program id, mint id])
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetNewMetadataAccount(newMetadata ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(newMetadata).WRITE()
	return inst
}

// GetNewMetadataAccount gets the "newMetadata" account.
// New Metadata key (pda of ['metadata', program id, mint id])
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetNewMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetNewEditionAccount sets the "newEdition" account.
// New Edition (pda of ['metadata', program id, mint id, 'edition'])
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetNewEditionAccount(newEdition ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(newEdition).WRITE()
	return inst
}

// GetNewEditionAccount gets the "newEdition" account.
// New Edition (pda of ['metadata', program id, mint id, 'edition'])
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetNewEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMasterEditionAccount sets the "masterEdition" account.
// Master Record Edition V2 (pda of ['metadata', program id, master metadata mint id, 'edition']
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
// Master Record Edition V2 (pda of ['metadata', program id, master metadata mint id, 'edition']
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetNewMintAccount sets the "newMint" account.
// Mint of new token - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetNewMintAccount(newMint ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(newMint).WRITE()
	return inst
}

// GetNewMintAccount gets the "newMint" account.
// Mint of new token - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetNewMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEditionMarkPdaAccount sets the "editionMarkPda" account.
// Edition pda to mark creation - will be checked for pre-existence. (pda of ['metadata', program id, master metadata mint id, 'edition', edition_number]) where edition_number is NOT the edition number you pass in args but actually edition_number = floor(edition/EDITION_MARKER_BIT_SIZE).
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetEditionMarkPdaAccount(editionMarkPda ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(editionMarkPda).WRITE()
	return inst
}

// GetEditionMarkPdaAccount gets the "editionMarkPda" account.
// Edition pda to mark creation - will be checked for pre-existence. (pda of ['metadata', program id, master metadata mint id, 'edition', edition_number]) where edition_number is NOT the edition number you pass in args but actually edition_number = floor(edition/EDITION_MARKER_BIT_SIZE).
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetEditionMarkPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetNewMintAuthorityAccount sets the "newMintAuthority" account.
// Mint authority of new mint
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetNewMintAuthorityAccount(newMintAuthority ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(newMintAuthority).SIGNER()
	return inst
}

// GetNewMintAuthorityAccount gets the "newMintAuthority" account.
// Mint authority of new mint
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetNewMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetPayerAccount(payer ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Vault authority
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(vaultAuthority).SIGNER()
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Vault authority
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSafetyDepositStoreAccount sets the "safetyDepositStore" account.
// Safety deposit token store account
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetSafetyDepositStoreAccount(safetyDepositStore ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(safetyDepositStore)
	return inst
}

// GetSafetyDepositStoreAccount gets the "safetyDepositStore" account.
// Safety deposit token store account
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetSafetyDepositStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// Safety deposit box
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(safetyDepositBox)
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// Safety deposit box
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetVaultAccount sets the "vault" account.
// Vault
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetVaultAccount(vault ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(vault)
	return inst
}

// GetVaultAccount gets the "vault" account.
// Vault
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetNewMetadataUpdateAuthorityAccount sets the "newMetadataUpdateAuthority" account.
// Update authority info for new metadata
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetNewMetadataUpdateAuthorityAccount(newMetadataUpdateAuthority ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(newMetadataUpdateAuthority)
	return inst
}

// GetNewMetadataUpdateAuthorityAccount gets the "newMetadataUpdateAuthority" account.
// Update authority info for new metadata
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetNewMetadataUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetMetadataAccount sets the "metadata" account.
// Master record metadata account
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetMetadataAccount(metadata ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Master record metadata account
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTokenVaultProgramAccount sets the "tokenVaultProgram" account.
// Token vault program
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetTokenVaultProgramAccount(tokenVaultProgram ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenVaultProgram)
	return inst
}

// GetTokenVaultProgramAccount gets the "tokenVaultProgram" account.
// Token vault program
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetTokenVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) SetRentAccount(rent ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account (optional).
// Rent info
func (inst *MintNewEditionFromMasterEditionViaVaultProxy) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

func (inst MintNewEditionFromMasterEditionViaVaultProxy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_MintNewEditionFromMasterEditionViaVaultProxy),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintNewEditionFromMasterEditionViaVaultProxy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintNewEditionFromMasterEditionViaVaultProxy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MintNewEditionFromMasterEditionViaTokenArgs == nil {
			return errors.New("MintNewEditionFromMasterEditionViaTokenArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.NewMetadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.NewEdition is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.NewMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EditionMarkPda is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.NewMintAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.VaultAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SafetyDepositStore is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.NewMetadataUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenVaultProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}

		// [16] = Rent is optional

	}
	return nil
}

func (inst *MintNewEditionFromMasterEditionViaVaultProxy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintNewEditionFromMasterEditionViaVaultProxy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MintNewEditionFromMasterEditionViaTokenArgs", *inst.MintNewEditionFromMasterEditionViaTokenArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=17]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               newMetadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                newEdition", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             masterEdition", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                   newMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            editionMarkPda", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          newMintAuthority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                     payer", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            vaultAuthority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("        safetyDepositStore", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          safetyDepositBox", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                     vault", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("newMetadataUpdateAuthority", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                  metadata", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("              tokenProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         tokenVaultProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("             systemProgram", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("                      rent", inst.AccountMetaSlice.Get(16)))
					})
				})
		})
}

func (obj MintNewEditionFromMasterEditionViaVaultProxy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MintNewEditionFromMasterEditionViaTokenArgs` param:
	err = encoder.Encode(obj.MintNewEditionFromMasterEditionViaTokenArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MintNewEditionFromMasterEditionViaVaultProxy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MintNewEditionFromMasterEditionViaTokenArgs`:
	err = decoder.Decode(&obj.MintNewEditionFromMasterEditionViaTokenArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewMintNewEditionFromMasterEditionViaVaultProxyInstruction declares a new MintNewEditionFromMasterEditionViaVaultProxy instruction with the provided parameters and accounts.
func NewMintNewEditionFromMasterEditionViaVaultProxyInstruction(
	// Parameters:
	mintNewEditionFromMasterEditionViaTokenArgs MintNewEditionFromMasterEditionViaTokenArgs,
	// Accounts:
	newMetadata ag_solanago.PublicKey,
	newEdition ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	newMint ag_solanago.PublicKey,
	editionMarkPda ag_solanago.PublicKey,
	newMintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	safetyDepositStore ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	newMetadataUpdateAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenVaultProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *MintNewEditionFromMasterEditionViaVaultProxy {
	return NewMintNewEditionFromMasterEditionViaVaultProxyInstructionBuilder().
		SetMintNewEditionFromMasterEditionViaTokenArgs(mintNewEditionFromMasterEditionViaTokenArgs).
		SetNewMetadataAccount(newMetadata).
		SetNewEditionAccount(newEdition).
		SetMasterEditionAccount(masterEdition).
		SetNewMintAccount(newMint).
		SetEditionMarkPdaAccount(editionMarkPda).
		SetNewMintAuthorityAccount(newMintAuthority).
		SetPayerAccount(payer).
		SetVaultAuthorityAccount(vaultAuthority).
		SetSafetyDepositStoreAccount(safetyDepositStore).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetVaultAccount(vault).
		SetNewMetadataUpdateAuthorityAccount(newMetadataUpdateAuthority).
		SetMetadataAccount(metadata).
		SetTokenProgramAccount(tokenProgram).
		SetTokenVaultProgramAccount(tokenVaultProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
