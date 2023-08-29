package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontDownloadOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.4 18h2.1l1.1-3.05h4.8L15.5 18h2.1L13.05 6h-2.1L6.4 18Zm3.8-4.8l1.75-4.95h.1l1.75 4.95h-3.6ZM2 22V2h20v20H2Zm2-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}