// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package power

import (
	"fmt"
	"io"
	"math"
	"sort"

	address "github.com/filecoin-project/go-address"
	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufState = []byte{143}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufState); err != nil {
		return err
	}

	// t.TotalRawBytePower (big.Int) (struct)
	if err := t.TotalRawBytePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TotalBytesCommitted (big.Int) (struct)
	if err := t.TotalBytesCommitted.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TotalQualityAdjPower (big.Int) (struct)
	if err := t.TotalQualityAdjPower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TotalQABytesCommitted (big.Int) (struct)
	if err := t.TotalQABytesCommitted.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TotalPledgeCollateral (big.Int) (struct)
	if err := t.TotalPledgeCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochRawBytePower (big.Int) (struct)
	if err := t.ThisEpochRawBytePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochQualityAdjPower (big.Int) (struct)
	if err := t.ThisEpochQualityAdjPower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochPledgeCollateral (big.Int) (struct)
	if err := t.ThisEpochPledgeCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochQAPowerSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.ThisEpochQAPowerSmoothed.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.MinerCount (int64) (int64)
	if t.MinerCount >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.MinerCount)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.MinerCount-1)); err != nil {
			return err
		}
	}

	// t.MinerAboveMinPowerCount (int64) (int64)
	if t.MinerAboveMinPowerCount >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.MinerAboveMinPowerCount)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.MinerAboveMinPowerCount-1)); err != nil {
			return err
		}
	}

	// t.CronEventQueue (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.CronEventQueue); err != nil {
		return xerrors.Errorf("failed to write cid field t.CronEventQueue: %w", err)
	}

	// t.FirstCronEpoch (abi.ChainEpoch) (int64)
	if t.FirstCronEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.FirstCronEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.FirstCronEpoch-1)); err != nil {
			return err
		}
	}

	// t.Claims (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.Claims); err != nil {
		return xerrors.Errorf("failed to write cid field t.Claims: %w", err)
	}

	// t.ProofValidationBatch (cid.Cid) (struct)

	if t.ProofValidationBatch == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.ProofValidationBatch); err != nil {
			return xerrors.Errorf("failed to write cid field t.ProofValidationBatch: %w", err)
		}
	}

	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) (err error) {
	*t = State{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 15 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 15", extra)
	}

	// t.TotalRawBytePower (big.Int) (struct)

	{

		if err := t.TotalRawBytePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalRawBytePower: %w", err)
		}

	}
	// t.TotalBytesCommitted (big.Int) (struct)

	{

		if err := t.TotalBytesCommitted.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalBytesCommitted: %w", err)
		}

	}
	// t.TotalQualityAdjPower (big.Int) (struct)

	{

		if err := t.TotalQualityAdjPower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalQualityAdjPower: %w", err)
		}

	}
	// t.TotalQABytesCommitted (big.Int) (struct)

	{

		if err := t.TotalQABytesCommitted.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalQABytesCommitted: %w", err)
		}

	}
	// t.TotalPledgeCollateral (big.Int) (struct)

	{

		if err := t.TotalPledgeCollateral.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalPledgeCollateral: %w", err)
		}

	}
	// t.ThisEpochRawBytePower (big.Int) (struct)

	{

		if err := t.ThisEpochRawBytePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochRawBytePower: %w", err)
		}

	}
	// t.ThisEpochQualityAdjPower (big.Int) (struct)

	{

		if err := t.ThisEpochQualityAdjPower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochQualityAdjPower: %w", err)
		}

	}
	// t.ThisEpochPledgeCollateral (big.Int) (struct)

	{

		if err := t.ThisEpochPledgeCollateral.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochPledgeCollateral: %w", err)
		}

	}
	// t.ThisEpochQAPowerSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.ThisEpochQAPowerSmoothed.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochQAPowerSmoothed: %w", err)
		}

	}
	// t.MinerCount (int64) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.MinerCount = int64(extraI)
	}
	// t.MinerAboveMinPowerCount (int64) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.MinerAboveMinPowerCount = int64(extraI)
	}
	// t.CronEventQueue (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.CronEventQueue: %w", err)
		}

		t.CronEventQueue = c

	}
	// t.FirstCronEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.FirstCronEpoch = abi.ChainEpoch(extraI)
	}
	// t.Claims (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Claims: %w", err)
		}

		t.Claims = c

	}
	// t.ProofValidationBatch (cid.Cid) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(cr)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.ProofValidationBatch: %w", err)
			}

			t.ProofValidationBatch = &c
		}

	}
	return nil
}

var lengthBufClaim = []byte{131}

func (t *Claim) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufClaim); err != nil {
		return err
	}

	// t.WindowPoStProofType (abi.RegisteredPoStProof) (int64)
	if t.WindowPoStProofType >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.WindowPoStProofType)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.WindowPoStProofType-1)); err != nil {
			return err
		}
	}

	// t.RawBytePower (big.Int) (struct)
	if err := t.RawBytePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.QualityAdjPower (big.Int) (struct)
	if err := t.QualityAdjPower.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *Claim) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Claim{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 3", extra)
	}

	// t.WindowPoStProofType (abi.RegisteredPoStProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.WindowPoStProofType = abi.RegisteredPoStProof(extraI)
	}
	// t.RawBytePower (big.Int) (struct)

	{

		if err := t.RawBytePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.RawBytePower: %w", err)
		}

	}
	// t.QualityAdjPower (big.Int) (struct)

	{

		if err := t.QualityAdjPower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.QualityAdjPower: %w", err)
		}

	}
	return nil
}

var lengthBufUpdateClaimedPowerParams = []byte{130}

func (t *UpdateClaimedPowerParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufUpdateClaimedPowerParams); err != nil {
		return err
	}

	// t.RawByteDelta (big.Int) (struct)
	if err := t.RawByteDelta.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.QualityAdjustedDelta (big.Int) (struct)
	if err := t.QualityAdjustedDelta.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *UpdateClaimedPowerParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = UpdateClaimedPowerParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 2", extra)
	}

	// t.RawByteDelta (big.Int) (struct)

	{

		if err := t.RawByteDelta.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.RawByteDelta: %w", err)
		}

	}
	// t.QualityAdjustedDelta (big.Int) (struct)

	{

		if err := t.QualityAdjustedDelta.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.QualityAdjustedDelta: %w", err)
		}

	}
	return nil
}

var lengthBufMinerConstructorParams = []byte{134}

func (t *MinerConstructorParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufMinerConstructorParams); err != nil {
		return err
	}

	// t.OwnerAddr (address.Address) (struct)
	if err := t.OwnerAddr.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.WorkerAddr (address.Address) (struct)
	if err := t.WorkerAddr.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ControlAddrs ([]address.Address) (slice)
	if len(t.ControlAddrs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.ControlAddrs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.ControlAddrs))); err != nil {
		return err
	}
	for _, v := range t.ControlAddrs {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}

	// t.WindowPoStProofType (abi.RegisteredPoStProof) (int64)
	if t.WindowPoStProofType >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.WindowPoStProofType)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.WindowPoStProofType-1)); err != nil {
			return err
		}
	}

	// t.PeerId ([]uint8) (slice)
	if len(t.PeerId) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.PeerId was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.PeerId))); err != nil {
		return err
	}

	if _, err := cw.Write(t.PeerId[:]); err != nil {
		return err
	}

	// t.Multiaddrs ([][]uint8) (slice)
	if len(t.Multiaddrs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Multiaddrs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Multiaddrs))); err != nil {
		return err
	}
	for _, v := range t.Multiaddrs {
		if len(v) > cbg.ByteArrayMaxLen {
			return xerrors.Errorf("Byte array in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(v))); err != nil {
			return err
		}

		if _, err := cw.Write(v[:]); err != nil {
			return err
		}
	}
	return nil
}

func (t *MinerConstructorParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MinerConstructorParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 6", extra)
	}

	// t.OwnerAddr (address.Address) (struct)

	{

		if err := t.OwnerAddr.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.OwnerAddr: %w", err)
		}

	}
	// t.WorkerAddr (address.Address) (struct)

	{

		if err := t.WorkerAddr.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.WorkerAddr: %w", err)
		}

	}
	// t.ControlAddrs ([]address.Address) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.ControlAddrs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.ControlAddrs = make([]address.Address, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				if err := t.ControlAddrs[i].UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ControlAddrs[i]: %w", err)
				}

			}
		}
	}

	// t.WindowPoStProofType (abi.RegisteredPoStProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.WindowPoStProofType = abi.RegisteredPoStProof(extraI)
	}
	// t.PeerId ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.PeerId: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.PeerId = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.PeerId[:]); err != nil {
		return err
	}
	// t.Multiaddrs ([][]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Multiaddrs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Multiaddrs = make([][]uint8, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Multiaddrs[i]: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Multiaddrs[i] = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Multiaddrs[i][:]); err != nil {
				return err
			}
		}
	}

	return nil
}

var lengthBufCreateMinerParams = []byte{133}

func (t *CreateMinerParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCreateMinerParams); err != nil {
		return err
	}

	// t.Owner (address.Address) (struct)
	if err := t.Owner.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Worker (address.Address) (struct)
	if err := t.Worker.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.WindowPoStProofType (abi.RegisteredPoStProof) (int64)
	if t.WindowPoStProofType >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.WindowPoStProofType)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.WindowPoStProofType-1)); err != nil {
			return err
		}
	}

	// t.Peer ([]uint8) (slice)
	if len(t.Peer) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Peer was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Peer))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Peer[:]); err != nil {
		return err
	}

	// t.Multiaddrs ([][]uint8) (slice)
	if len(t.Multiaddrs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Multiaddrs was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Multiaddrs))); err != nil {
		return err
	}
	for _, v := range t.Multiaddrs {
		if len(v) > cbg.ByteArrayMaxLen {
			return xerrors.Errorf("Byte array in field v was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(v))); err != nil {
			return err
		}

		if _, err := cw.Write(v[:]); err != nil {
			return err
		}
	}
	return nil
}

func (t *CreateMinerParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CreateMinerParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 5", extra)
	}

	// t.Owner (address.Address) (struct)

	{

		if err := t.Owner.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Owner: %w", err)
		}

	}
	// t.Worker (address.Address) (struct)

	{

		if err := t.Worker.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Worker: %w", err)
		}

	}
	// t.WindowPoStProofType (abi.RegisteredPoStProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.WindowPoStProofType = abi.RegisteredPoStProof(extraI)
	}
	// t.Peer ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Peer: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Peer = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Peer[:]); err != nil {
		return err
	}
	// t.Multiaddrs ([][]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Multiaddrs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Multiaddrs = make([][]uint8, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Multiaddrs[i]: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Multiaddrs[i] = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Multiaddrs[i][:]); err != nil {
				return err
			}
		}
	}

	return nil
}

var lengthBufCreateMinerReturn = []byte{130}

func (t *CreateMinerReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCreateMinerReturn); err != nil {
		return err
	}

	// t.IDAddress (address.Address) (struct)
	if err := t.IDAddress.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.RobustAddress (address.Address) (struct)
	if err := t.RobustAddress.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *CreateMinerReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CreateMinerReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 2", extra)
	}

	// t.IDAddress (address.Address) (struct)

	{

		if err := t.IDAddress.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.IDAddress: %w", err)
		}

	}
	// t.RobustAddress (address.Address) (struct)

	{

		if err := t.RobustAddress.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.RobustAddress: %w", err)
		}

	}
	return nil
}

var lengthBufCurrentTotalPowerReturn = []byte{132}

func (t *CurrentTotalPowerReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCurrentTotalPowerReturn); err != nil {
		return err
	}

	// t.RawBytePower (big.Int) (struct)
	if err := t.RawBytePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.QualityAdjPower (big.Int) (struct)
	if err := t.QualityAdjPower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.PledgeCollateral (big.Int) (struct)
	if err := t.PledgeCollateral.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.QualityAdjPowerSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.QualityAdjPowerSmoothed.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *CurrentTotalPowerReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CurrentTotalPowerReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 4", extra)
	}

	// t.RawBytePower (big.Int) (struct)

	{

		if err := t.RawBytePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.RawBytePower: %w", err)
		}

	}
	// t.QualityAdjPower (big.Int) (struct)

	{

		if err := t.QualityAdjPower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.QualityAdjPower: %w", err)
		}

	}
	// t.PledgeCollateral (big.Int) (struct)

	{

		if err := t.PledgeCollateral.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.PledgeCollateral: %w", err)
		}

	}
	// t.QualityAdjPowerSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.QualityAdjPowerSmoothed.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.QualityAdjPowerSmoothed: %w", err)
		}

	}
	return nil
}

var lengthBufEnrollCronEventParams = []byte{130}

func (t *EnrollCronEventParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufEnrollCronEventParams); err != nil {
		return err
	}

	// t.EventEpoch (abi.ChainEpoch) (int64)
	if t.EventEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.EventEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.EventEpoch-1)); err != nil {
			return err
		}
	}

	// t.Payload ([]uint8) (slice)
	if len(t.Payload) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Payload was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Payload))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Payload[:]); err != nil {
		return err
	}
	return nil
}

func (t *EnrollCronEventParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = EnrollCronEventParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 2", extra)
	}

	// t.EventEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.EventEpoch = abi.ChainEpoch(extraI)
	}
	// t.Payload ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Payload: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Payload = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Payload[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufMinerRawPowerReturn = []byte{130}

func (t *MinerRawPowerReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufMinerRawPowerReturn); err != nil {
		return err
	}

	// t.RawBytePower (big.Int) (struct)
	if err := t.RawBytePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.MeetsConsensusMinimum (bool) (bool)
	if err := cbg.WriteBool(w, t.MeetsConsensusMinimum); err != nil {
		return err
	}
	return nil
}

func (t *MinerRawPowerReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MinerRawPowerReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 2", extra)
	}

	// t.RawBytePower (big.Int) (struct)

	{

		if err := t.RawBytePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.RawBytePower: %w", err)
		}

	}
	// t.MeetsConsensusMinimum (bool) (bool)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.MeetsConsensusMinimum = false
	case 21:
		t.MeetsConsensusMinimum = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufCronEvent = []byte{130}

func (t *CronEvent) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufCronEvent); err != nil {
		return err
	}

	// t.MinerAddr (address.Address) (struct)
	if err := t.MinerAddr.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.CallbackPayload ([]uint8) (slice)
	if len(t.CallbackPayload) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.CallbackPayload was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.CallbackPayload))); err != nil {
		return err
	}

	if _, err := cw.Write(t.CallbackPayload[:]); err != nil {
		return err
	}
	return nil
}

func (t *CronEvent) UnmarshalCBOR(r io.Reader) (err error) {
	*t = CronEvent{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields, got %d, expected 2", extra)
	}

	// t.MinerAddr (address.Address) (struct)

	{

		if err := t.MinerAddr.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.MinerAddr: %w", err)
		}

	}
	// t.CallbackPayload ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.CallbackPayload: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.CallbackPayload = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.CallbackPayload[:]); err != nil {
		return err
	}
	return nil
}
