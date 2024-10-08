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

// Transfer is the `Transfer` instruction.
type Transfer struct {
	TransferArgs TransferArgs

	// [0] = [WRITE] token
	// ··········· Token account
	//
	// [1] = [] tokenOwner
	// ··········· Token account owner
	//
	// [2] = [WRITE] destination
	// ··········· Destination token account
	//
	// [3] = [] destinationOwner
	// ··········· Destination token account owner
	//
	// [4] = [] mint
	// ··········· Mint of token asset
	//
	// [5] = [WRITE] metadata
	// ··········· Metadata (pda of ['metadata', program id, mint id])
	//
	// [6] = [] edition
	// ··········· Edition of token asset
	//
	// [7] = [WRITE] ownerTokenRecord
	// ··········· Owner token record account
	//
	// [8] = [WRITE] destinationTokenRecord
	// ··········· Destination token record account
	//
	// [9] = [SIGNER] authority
	// ··········· Transfer authority (token owner or delegate)
	//
	// [10] = [WRITE, SIGNER] payer
	// ··········· Payer
	//
	// [11] = [] systemProgram
	// ··········· System Program
	//
	// [12] = [] sysvarInstructions
	// ··········· Instructions sysvar account
	//
	// [13] = [] splTokenProgram
	// ··········· SPL Token Program
	//
	// [14] = [] splAtaProgram
	// ··········· SPL Associated Token Account program
	//
	// [15] = [] authorizationRulesProgram
	// ··········· Token Authorization Rules Program
	//
	// [16] = [] authorizationRules
	// ··········· Token Authorization Rules account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferInstructionBuilder creates a new `Transfer` instruction builder.
func NewTransferInstructionBuilder() *Transfer {
	nd := &Transfer{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 17),
	}
	return nd
}

// SetTransferArgs sets the "transferArgs" parameter.
func (inst *Transfer) SetTransferArgs(transferArgs TransferArgs) *Transfer {
	inst.TransferArgs = transferArgs
	return inst
}

// SetTokenAccount sets the "token" account.
// Token account
func (inst *Transfer) SetTokenAccount(token ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(token).WRITE()
	return inst
}

// GetTokenAccount gets the "token" account.
// Token account
func (inst *Transfer) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenOwnerAccount sets the "tokenOwner" account.
// Token account owner
func (inst *Transfer) SetTokenOwnerAccount(tokenOwner ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenOwner)
	return inst
}

// GetTokenOwnerAccount gets the "tokenOwner" account.
// Token account owner
func (inst *Transfer) GetTokenOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDestinationAccount sets the "destination" account.
// Destination token account
func (inst *Transfer) SetDestinationAccount(destination ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
// Destination token account
func (inst *Transfer) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetDestinationOwnerAccount sets the "destinationOwner" account.
// Destination token account owner
func (inst *Transfer) SetDestinationOwnerAccount(destinationOwner ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(destinationOwner)
	return inst
}

// GetDestinationOwnerAccount gets the "destinationOwner" account.
// Destination token account owner
func (inst *Transfer) GetDestinationOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintAccount sets the "mint" account.
// Mint of token asset
func (inst *Transfer) SetMintAccount(mint ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// Mint of token asset
func (inst *Transfer) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata (pda of ['metadata', program id, mint id])
func (inst *Transfer) SetMetadataAccount(metadata ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata (pda of ['metadata', program id, mint id])
func (inst *Transfer) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEditionAccount sets the "edition" account.
// Edition of token asset
func (inst *Transfer) SetEditionAccount(edition ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(edition)
	return inst
}

// GetEditionAccount gets the "edition" account (optional).
// Edition of token asset
func (inst *Transfer) GetEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetOwnerTokenRecordAccount sets the "ownerTokenRecord" account.
// Owner token record account
func (inst *Transfer) SetOwnerTokenRecordAccount(ownerTokenRecord ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(ownerTokenRecord).WRITE()
	return inst
}

// GetOwnerTokenRecordAccount gets the "ownerTokenRecord" account (optional).
// Owner token record account
func (inst *Transfer) GetOwnerTokenRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDestinationTokenRecordAccount sets the "destinationTokenRecord" account.
// Destination token record account
func (inst *Transfer) SetDestinationTokenRecordAccount(destinationTokenRecord ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(destinationTokenRecord).WRITE()
	return inst
}

// GetDestinationTokenRecordAccount gets the "destinationTokenRecord" account (optional).
// Destination token record account
func (inst *Transfer) GetDestinationTokenRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAuthorityAccount sets the "authority" account.
// Transfer authority (token owner or delegate)
func (inst *Transfer) SetAuthorityAccount(authority ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Transfer authority (token owner or delegate)
func (inst *Transfer) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *Transfer) SetPayerAccount(payer ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *Transfer) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System Program
func (inst *Transfer) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System Program
func (inst *Transfer) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *Transfer) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *Transfer) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// SPL Token Program
func (inst *Transfer) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account.
// SPL Token Program
func (inst *Transfer) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSplAtaProgramAccount sets the "splAtaProgram" account.
// SPL Associated Token Account program
func (inst *Transfer) SetSplAtaProgramAccount(splAtaProgram ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(splAtaProgram)
	return inst
}

// GetSplAtaProgramAccount gets the "splAtaProgram" account.
// SPL Associated Token Account program
func (inst *Transfer) GetSplAtaProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetAuthorizationRulesProgramAccount sets the "authorizationRulesProgram" account.
// Token Authorization Rules Program
func (inst *Transfer) SetAuthorizationRulesProgramAccount(authorizationRulesProgram ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(authorizationRulesProgram)
	return inst
}

// GetAuthorizationRulesProgramAccount gets the "authorizationRulesProgram" account (optional).
// Token Authorization Rules Program
func (inst *Transfer) GetAuthorizationRulesProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetAuthorizationRulesAccount sets the "authorizationRules" account.
// Token Authorization Rules account
func (inst *Transfer) SetAuthorizationRulesAccount(authorizationRules ag_solanago.PublicKey) *Transfer {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(authorizationRules)
	return inst
}

// GetAuthorizationRulesAccount gets the "authorizationRules" account (optional).
// Token Authorization Rules account
func (inst *Transfer) GetAuthorizationRulesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

func (inst Transfer) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_Transfer),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Transfer) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Transfer) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TransferArgs == nil {
			return errors.New("TransferArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Token is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenOwner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Destination is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.DestinationOwner is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}

		// [6] = Edition is optional

		// [7] = OwnerTokenRecord is optional

		// [8] = DestinationTokenRecord is optional

		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SplTokenProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SplAtaProgram is not set")
		}

		// [15] = AuthorizationRulesProgram is optional

		// [16] = AuthorizationRules is optional

	}
	return nil
}

func (inst *Transfer) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Transfer")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("TransferArgs", inst.TransferArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=17]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    token", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               tokenOwner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              destination", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         destinationOwner", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                     mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                  edition", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         ownerTokenRecord", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("   destinationTokenRecord", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("            systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("       sysvarInstructions", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("          splTokenProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("            splAtaProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("authorizationRulesProgram", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("       authorizationRules", inst.AccountMetaSlice.Get(16)))
					})
				})
		})
}

func (obj Transfer) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TransferArgs` param:
	{
		tmp := transferArgsContainer{}
		switch realvalue := obj.TransferArgs.(type) {
		case *TransferArgsV1:
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
func (obj *Transfer) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TransferArgs`:
	{
		tmp := new(transferArgsContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.TransferArgs = &tmp.V1
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	return nil
}

// NewTransferInstruction declares a new Transfer instruction with the provided parameters and accounts.
func NewTransferInstruction(
	// Parameters:
	transferArgs TransferArgs,
	// Accounts:
	token ag_solanago.PublicKey,
	tokenOwner ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	destinationOwner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	edition ag_solanago.PublicKey,
	ownerTokenRecord ag_solanago.PublicKey,
	destinationTokenRecord ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey,
	splAtaProgram ag_solanago.PublicKey,
	authorizationRulesProgram ag_solanago.PublicKey,
	authorizationRules ag_solanago.PublicKey) *Transfer {
	return NewTransferInstructionBuilder().
		SetTransferArgs(transferArgs).
		SetTokenAccount(token).
		SetTokenOwnerAccount(tokenOwner).
		SetDestinationAccount(destination).
		SetDestinationOwnerAccount(destinationOwner).
		SetMintAccount(mint).
		SetMetadataAccount(metadata).
		SetEditionAccount(edition).
		SetOwnerTokenRecordAccount(ownerTokenRecord).
		SetDestinationTokenRecordAccount(destinationTokenRecord).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetSplTokenProgramAccount(splTokenProgram).
		SetSplAtaProgramAccount(splAtaProgram).
		SetAuthorizationRulesProgramAccount(authorizationRulesProgram).
		SetAuthorizationRulesAccount(authorizationRules)
}
