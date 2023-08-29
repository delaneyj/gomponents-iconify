package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoomServiceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-2h20v2H2Zm1-3v-1q0-3.2 1.963-5.65T10 6.25V4h4v2.25q3.1.65 5.05 3.1T21 15v1H3Z"/>`),
		g.Group(children),
	)
}