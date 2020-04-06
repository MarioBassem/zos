package identity

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tyler-smith/go-bip39"

	"github.com/threefoldtech/zos/pkg/versioned"
)

// UserIdentity defines serializable struct to identify a user
type UserIdentity struct {
	// Mnemonic words of Private Key
	Mnemonic string `json:"mnemonic"`
	// ThreebotID generated by explorer
	ThreebotID uint64 `json:"threebotid"`
	// Internal keypair not exported
	key KeyPair
}

// NewUserIdentity create a new UserIdentity from existing key
func NewUserIdentity(key KeyPair, threebotid uint64) UserIdentity {
	return UserIdentity{
		key:        key,
		ThreebotID: threebotid,
	}
}

// Key returns the internal KeyPair
func (u *UserIdentity) Key() KeyPair {
	return u.key
}

// Load fetch a seed file and initialize key based on mnemonic
func (u *UserIdentity) Load(path string) error {
	version, buf, err := versioned.ReadFile(path)
	if err != nil {
		return err
	}

	if version.Compare(seedVersion1) == 0 {
		return fmt.Errorf("seed file too old, please update it using 'tfuser id convert' command")
	}

	if version.NE(seedVersionLatest) {
		return fmt.Errorf("unsupported seed version")
	}

	err = json.Unmarshal(buf, &u)
	if err != nil {
		return err
	}

	return u.FromMnemonic(u.Mnemonic)
}

// FromMnemonic initialize the Key (KeyPair) from mnemonic argument
func (u *UserIdentity) FromMnemonic(mnemonic string) error {
	seed, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	// Loading mnemonic
	u.key, err = FromSeed(seed)
	if err != nil {
		return err
	}

	return nil
}

// Save dumps UserIdentity into a versioned file
func (u *UserIdentity) Save(path string) error {
	var err error

	log.Info().Msg("generating seed mnemonic")

	// Generate mnemonic of private key
	u.Mnemonic, err = bip39.NewMnemonic(u.key.PrivateKey.Seed())
	if err != nil {
		return err
	}

	// Versioning json output
	buf, err := json.Marshal(u)
	if err != nil {
		return err
	}

	// Saving json to file
	log.Info().Str("filename", path).Msg("writing user identity")
	versioned.WriteFile(path, seedVersion11, buf, 0400)

	return nil
}
