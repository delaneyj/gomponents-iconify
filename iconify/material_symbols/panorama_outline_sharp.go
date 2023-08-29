package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h16V6H4v12Zm2-2h12l-3.75-5l-3 4L9 12l-3 4Zm-2 2V6v12Z"/>`),
		g.Group(children),
	)
}