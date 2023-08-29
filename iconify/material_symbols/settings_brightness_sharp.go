package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsBrightnessSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17.5l1.5-1.5H16v-2.5l1.5-1.5l-1.5-1.5V8h-2.5L12 6.5L10.5 8H8v2.5L6.5 12L8 13.5V16h2.5l1.5 1.5Zm0-2.5V9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}