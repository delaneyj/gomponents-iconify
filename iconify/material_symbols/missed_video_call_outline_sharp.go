package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MissedVideoCallOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.7 16l4.5-4.55l-1.4-1.4l-3.1 3.1L8.5 11H10V9H5v5h2v-1.7l3.7 3.7ZM2 20V4h16v6.5l4-4v11l-4-4V20H2Zm2-2h12V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}