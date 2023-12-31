// Copyright 2020 EMIT Foundation.
// This file is part of E.M.I.T. .

// E.M.I.T. is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// E.M.I.T. is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with E.M.I.T. . If not, see <http://www.gnu.org/licenses/>.

package ethereum

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (w *writer) GetDepositFTRecord(nonce uint64, dest uint8) (resourceID [32]byte, destinationRecipientAddress []byte, amount *big.Int, err error) {
	prop, err := w.erc20HandlerContract.GetDepositRecord(w.conn.CallOpts(), nonce, dest)
	if err != nil {
		return [32]byte{}, nil, nil, err
	}
	return prop.ResourceID, prop.DestinationRecipientAddress, prop.Amount, nil

}
func (w *writer) GetDepositNFTRecord(nonce uint64, dest uint8) (src721ResourceID [32]byte,
	destinationRecipientAddress []byte,
	tokenId *big.Int,
	metadata []byte,
	src20Amount *big.Int,
	err error) {
	prop, err := w.erc721HandlerContract.GetDepositRecord(w.conn.CallOpts(), nonce, dest)
	if err != nil {
		return [32]byte{}, nil, nil, nil, nil, err
	}
	return prop.ResourceID, prop.DestinationRecipientAddress, prop.TokenID, prop.MetaData, nil, nil

}
func (w *writer) FTPropsalDataHash(recipient []byte, amount *big.Int) [32]byte {
	return ConstructErc20ProposalDataHash(w.cfg.erc20HandlerContract, ethCommon.BytesToAddress(recipient), amount)
}
func (w *writer) GetFTProposalStatus(source uint8, nonce uint64, dataHash [32]byte) (uint8, error) {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(source), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return 0, err
	}
	return prop.Status, nil
}

func (w *writer) NFTPropsalDataHash(recipient []byte, amount *big.Int, metadata []byte, feeAmount *big.Int) [32]byte {
	return ConstructErc721ProposalDataHash(w.cfg.erc721HandlerContract, ethCommon.BytesToAddress(recipient), amount, metadata, feeAmount)
}
func (w *writer) GetNFTProposalStatus(source uint8, nonce uint64, dataHash [32]byte) (uint8, error) {
	prop, err := w.nftBridgeContract.GetProposal(w.conn.CallOpts(), uint8(source), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return 0, err
	}
	return prop.Status, nil
}

func (w *writer) GetBridgeAddress() []byte {
	return w.cfg.bridgeContract.Bytes()
}
func (w *writer) GetNFTBridgeAddress() []byte {
	return w.cfg.nftBridgeContract.Bytes()
}

func (w *writer) IsWithCollector() bool {
	return true
}
