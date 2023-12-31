// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/core"
	"github.com/emit-technology/emit-cross/types"
	"math/big"

	utils "github.com/emit-technology/emit-cross/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 100000000000
const DefaultBlockConfirmations = 12

// Chain specific options
var (
	BridgeOpt              = "bridge"
	NFTBridgeOpt           = "nftBridge"
	Erc20HandlerOpt        = "erc20Handler"
	Erc721HandlerOpt       = "erc721Handler"
	GenericHandlerOpt      = "genericHandler"
	MaxGasPriceOpt         = "maxGasPrice"
	GasLimitOpt            = "gasLimit"
	HttpOpt                = "http"
	StartBlockOpt          = "startBlock"
	commitNodeOpt          = "commitNode"
	BlockConfirmationsOpt  = "blockConfirmations"
	CommitVotesStartSeqOpt = "commitVotesStartSeq"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                  string        // Human-readable chain name
	id                    types.ChainId // ChainID
	endpoint              string        // url for rpc endpoint
	from                  string        // address of key to use
	keystorePath          string        // Location of keyfiles
	blockstorePath        string
	freshStart            bool // Disables loading from blockstore at start
	bridgeContract        common.Address
	nftBridgeContract     common.Address
	erc20HandlerContract  common.Address
	erc721HandlerContract common.Address
	gasLimit              *big.Int
	maxGasPrice           *big.Int
	http                  bool // Config for type of connection
	startBlock            *big.Int
	blockConfirmations    *big.Int
	commitVotesStartSeq   *big.Int
	commitNode            bool
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:                 chainCfg.Name,
		id:                   chainCfg.Id,
		endpoint:             chainCfg.Endpoint,
		from:                 chainCfg.From,
		keystorePath:         chainCfg.KeystorePath,
		freshStart:           chainCfg.FreshStart,
		bridgeContract:       utils.ZeroAddress,
		erc20HandlerContract: utils.ZeroAddress,
		gasLimit:             big.NewInt(DefaultGasLimit),
		maxGasPrice:          big.NewInt(DefaultGasPrice),
		http:                 false,
		startBlock:           big.NewInt(0),
		blockConfirmations:   big.NewInt(0),
		commitVotesStartSeq:  big.NewInt(0),
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.bridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for sero config")
	}

	if contract, ok := chainCfg.Opts[NFTBridgeOpt]; ok && contract != "" {
		config.nftBridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, NFTBridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.nftbridge field for sero config")
	}

	config.erc20HandlerContract = common.HexToAddress(chainCfg.Opts[Erc20HandlerOpt])
	delete(chainCfg.Opts, Erc20HandlerOpt)

	config.erc721HandlerContract = common.HexToAddress(chainCfg.Opts[Erc721HandlerOpt])
	delete(chainCfg.Opts, Erc721HandlerOpt)

	if gasPrice, ok := chainCfg.Opts[MaxGasPriceOpt]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.maxGasPrice = price
			delete(chainCfg.Opts, MaxGasPriceOpt)
		} else {
			return nil, errors.New("unable to parse max gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts[GasLimitOpt]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
			delete(chainCfg.Opts, GasLimitOpt)
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, HttpOpt)
	} else if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, HttpOpt)
	}

	if commit, ok := chainCfg.Opts[commitNodeOpt]; ok && commit == "true" {
		config.commitNode = true
		delete(chainCfg.Opts, commitNodeOpt)
	} else if commit, ok := chainCfg.Opts[commitNodeOpt]; ok && commit == "false" {
		config.commitNode = false
		delete(chainCfg.Opts, commitNodeOpt)
	}

	if startBlock, ok := chainCfg.Opts[StartBlockOpt]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.startBlock = block
			delete(chainCfg.Opts, StartBlockOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", StartBlockOpt)
		}
	}

	if commitvotesStartSeq, ok := chainCfg.Opts[CommitVotesStartSeqOpt]; ok && commitvotesStartSeq != "" {
		start := big.NewInt(0)
		_, pass := start.SetString(commitvotesStartSeq, 10)
		if pass {
			config.commitVotesStartSeq = start
			delete(chainCfg.Opts, CommitVotesStartSeqOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", CommitVotesStartSeqOpt)
		}
	}

	if blockConfirmations, ok := chainCfg.Opts[BlockConfirmationsOpt]; ok && blockConfirmations != "" {
		val := big.NewInt(DefaultBlockConfirmations)
		_, pass := val.SetString(blockConfirmations, 10)
		if pass {
			config.blockConfirmations = val
			delete(chainCfg.Opts, BlockConfirmationsOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", BlockConfirmationsOpt)
		}
	}

	if len(chainCfg.Opts) != 0 {
		return nil, fmt.Errorf("unknown Opts Encountered: %#v", chainCfg.Opts)
	}

	return config, nil
}
