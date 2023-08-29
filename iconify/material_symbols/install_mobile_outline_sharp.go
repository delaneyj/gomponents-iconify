package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallMobileOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23V1h9v2H7v1h7v2H7v12h10v-2h2v7H5Zm2-3v1h10v-1H7Zm11-6l-5-5l1.4-1.4l2.6 2.6V3h2v7.2l2.6-2.6L23 9l-5 5ZM7 4V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}