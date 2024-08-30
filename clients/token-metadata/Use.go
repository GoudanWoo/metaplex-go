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

// Use is the `Use` instruction.
type Use struct {
	UseArgs UseArgs

	// [0] = [SIGNER] authority
	// ··········· Token owner or delegate
	//
	// [1] = [WRITE] delegateRecord
	// ··········· Delegate record PDA
	//
	// [2] = [WRITE] token
	// ··········· Token account
	//
	// [3] = [] mint
	// ··········· Mint account
	//
	// [4] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [5] = [WRITE] edition
	// ··········· Edition account
	//
	// [6] = [SIGNER] payer
	// ··········· Payer
	//
	// [7] = [] systemProgram
	// ··········· System program
	//
	// [8] = [] sysvarInstructions
	// ··········· System program
	//
	// [9] = [] splTokenProgram
	// ··········· SPL Token Program
	//
	// [10] = [] authorizationRulesProgram
	// ··········· Token Authorization Rules Program
	//
	// [11] = [] authorizationRules
	// ··········· Token Authorization Rules account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUseInstructionBuilder creates a new `Use` instruction builder.
func NewUseInstructionBuilder() *Use {
	nd := &Use{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetUseArgs sets the "useArgs" parameter.
func (inst *Use) SetUseArgs(useArgs UseArgs) *Use {
	inst.UseArgs = useArgs
	return inst
}

// SetAuthorityAccount sets the "authority" account.
// Token owner or delegate
func (inst *Use) SetAuthorityAccount(authority ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Token owner or delegate
func (inst *Use) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetDelegateRecordAccount sets the "delegateRecord" account.
// Delegate record PDA
func (inst *Use) SetDelegateRecordAccount(delegateRecord ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(delegateRecord).WRITE()
	return inst
}

// GetDelegateRecordAccount gets the "delegateRecord" account (optional).
// Delegate record PDA
func (inst *Use) GetDelegateRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenAccount sets the "token" account.
// Token account
func (inst *Use) SetTokenAccount(token ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(token).WRITE()
	return inst
}

// GetTokenAccount gets the "token" account (optional).
// Token account
func (inst *Use) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
// Mint account
func (inst *Use) SetMintAccount(mint ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// Mint account
func (inst *Use) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *Use) SetMetadataAccount(metadata ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *Use) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetEditionAccount sets the "edition" account.
// Edition account
func (inst *Use) SetEditionAccount(edition ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(edition).WRITE()
	return inst
}

// GetEditionAccount gets the "edition" account (optional).
// Edition account
func (inst *Use) GetEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *Use) SetPayerAccount(payer ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *Use) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *Use) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *Use) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
// System program
func (inst *Use) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
// System program
func (inst *Use) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// SPL Token Program
func (inst *Use) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account (optional).
// SPL Token Program
func (inst *Use) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAuthorizationRulesProgramAccount sets the "authorizationRulesProgram" account.
// Token Authorization Rules Program
func (inst *Use) SetAuthorizationRulesProgramAccount(authorizationRulesProgram ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(authorizationRulesProgram)
	return inst
}

// GetAuthorizationRulesProgramAccount gets the "authorizationRulesProgram" account (optional).
// Token Authorization Rules Program
func (inst *Use) GetAuthorizationRulesProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAuthorizationRulesAccount sets the "authorizationRules" account.
// Token Authorization Rules account
func (inst *Use) SetAuthorizationRulesAccount(authorizationRules ag_solanago.PublicKey) *Use {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(authorizationRules)
	return inst
}

// GetAuthorizationRulesAccount gets the "authorizationRules" account (optional).
// Token Authorization Rules account
func (inst *Use) GetAuthorizationRulesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst Use) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_Use),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Use) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Use) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UseArgs == nil {
			return errors.New("UseArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}

		// [1] = DelegateRecord is optional

		// [2] = Token is optional

		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Metadata is not set")
		}

		// [5] = Edition is optional

		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}

		// [9] = SplTokenProgram is optional

		// [10] = AuthorizationRulesProgram is optional

		// [11] = AuthorizationRules is optional

	}
	return nil
}

func (inst *Use) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Use")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UseArgs", inst.UseArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           delegateRecord", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    token", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                     mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                  edition", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       sysvarInstructions", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          splTokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("authorizationRulesProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       authorizationRules", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj Use) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UseArgs` param:
	{
		tmp := useArgsContainer{}
		switch realvalue := obj.UseArgs.(type) {
		case *UseArgsV1:
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
func (obj *Use) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UseArgs`:
	{
		tmp := new(useArgsContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.UseArgs = &tmp.V1
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	return nil
}

// NewUseInstruction declares a new Use instruction with the provided parameters and accounts.
func NewUseInstruction(
	// Parameters:
	useArgs UseArgs,
	// Accounts:
	authority ag_solanago.PublicKey,
	delegateRecord ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	edition ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey,
	authorizationRulesProgram ag_solanago.PublicKey,
	authorizationRules ag_solanago.PublicKey) *Use {
	return NewUseInstructionBuilder().
		SetUseArgs(useArgs).
		SetAuthorityAccount(authority).
		SetDelegateRecordAccount(delegateRecord).
		SetTokenAccount(token).
		SetMintAccount(mint).
		SetMetadataAccount(metadata).
		SetEditionAccount(edition).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetSplTokenProgramAccount(splTokenProgram).
		SetAuthorizationRulesProgramAccount(authorizationRulesProgram).
		SetAuthorizationRulesAccount(authorizationRules)
}
