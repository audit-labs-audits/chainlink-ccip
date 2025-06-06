// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_receiver

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
)

type Counter struct {
	Value     uint64
	RejectAll bool
	State     BaseState
}

var CounterDiscriminator = [8]byte{255, 176, 4, 245, 188, 253, 124, 25}

func (obj Counter) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CounterDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Value` param:
	err = encoder.Encode(obj.Value)
	if err != nil {
		return err
	}
	// Serialize `RejectAll` param:
	err = encoder.Encode(obj.RejectAll)
	if err != nil {
		return err
	}
	// Serialize `State` param:
	err = encoder.Encode(obj.State)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Counter) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CounterDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[255 176 4 245 188 253 124 25]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Value`:
	err = decoder.Decode(&obj.Value)
	if err != nil {
		return err
	}
	// Deserialize `RejectAll`:
	err = decoder.Decode(&obj.RejectAll)
	if err != nil {
		return err
	}
	// Deserialize `State`:
	err = decoder.Decode(&obj.State)
	if err != nil {
		return err
	}
	return nil
}

type ExternalExecutionConfig struct{}

var ExternalExecutionConfigDiscriminator = [8]byte{159, 157, 150, 212, 168, 103, 117, 39}

func (obj ExternalExecutionConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ExternalExecutionConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ExternalExecutionConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ExternalExecutionConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[159 157 150 212 168 103 117 39]",
				fmt.Sprint(discriminator[:]))
		}
	}
	return nil
}
