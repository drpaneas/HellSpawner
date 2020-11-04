package hsconfig

type AppConfig struct {
	Colors AppColorConfig `json:"colors"`
	Fonts  FontConfig     `json:"fonts"`
}

type AppColorConfig struct {
	WindowBackground     []uint8 `json:"windowBackground"`
	Text                 []uint8 `json:"text"`
	Primary              []uint8 `json:"primary"`
	PrimaryHighlight     []uint8 `json:"primaryHighlight"`
	Disabled             []uint8 `json:"disabled"`
	DisabledText         []uint8 `json:"disabledText"`
	Tab                  []uint8 `json:"tab"`
	TabSelected          []uint8 `json:"tabSelected"`
	TabHighlight         []uint8 `json:"tabHighlight"`
	TabSelectedHighlight []uint8 `json:"tabSelectedHighlight"`
}

type FontConfig struct {
	Normal     FontItemConfig `json:"normal"`
	Symbols    FontItemConfig `json:"symbols"`
	Monospaced FontItemConfig `json:"monospaced"`
	Info       FontItemConfig `json:"info"`
}

type FontItemConfig struct {
	Face string `json:"face"`
	Size int    `json:"size"`
}
