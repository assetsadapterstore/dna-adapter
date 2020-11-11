/*
 * Copyright 2020 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package dna

import (
	"github.com/blocktree/bitshares-adapter/v2/bitshares"
	"github.com/blocktree/openwallet/v2/log"
	bts "github.com/denkhaus/bitshares"
	"github.com/denkhaus/bitshares/config"
)

const (
	ChainIDDNA = "24938a99198d850bb7d79010c1325fb63fde63e8e477a5443ff5ce50ab867055"
)

type WalletManager struct {
	*bitshares.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = bitshares.NewWalletManager(nil)
	wm.Config = bitshares.NewConfig(Symbol)
	wm.Decoder = NewAddressDecoder(&wm)
	wm.DecoderV2 = NewAddressDecoder(&wm)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.Api = bitshares.NewWalletClient(wm.Config.ServerAPI, wm.Config.WalletAPI, false)
	wm.WebsocketAPI = NewWebsocketAPI(wm.Config.ServerWS)
	return &wm
}

func NewWebsocketAPI(api string) bts.WebsocketAPI {
	config.Add(config.ChainConfig{
		Name:      "DNA",
		CoreAsset: "DNA",
		Prefix:    "DNA",
		ID:        ChainIDDNA,
	})
	config.SetCurrent(ChainIDDNA)
	return bts.NewWebsocketAPI(api)
}
