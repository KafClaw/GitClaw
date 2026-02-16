// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

const (
	forcedI18nLang = "en-US"
	forcedI18nName = "English"
)

var (
	// I18n settings
	Langs []string
	Names []string
)

func loadI18nFrom(rootCfg ConfigProvider) {
	// GitClaw is English-only by design to keep agent and operator behavior deterministic.
	Langs = []string{forcedI18nLang}
	Names = []string{forcedI18nName}
}
