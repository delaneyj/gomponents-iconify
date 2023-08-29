package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenShareSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-2h22v2H1Zm1-3V3h20v15H2Zm6-4h2v-2q0-.425.288-.713T11 11h2v2l3-3l-3-3v2h-2q-1.25 0-2.125.875T8 12v2Z"/>`),
		g.Group(children),
	)
}