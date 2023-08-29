package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4h1V2H5v2h13Zm3 3.5H3v-2h18v2ZM23 9v13H1V9h22Zm-2 2H3v9h18v-9Z"/>`),
		g.Group(children),
	)
}