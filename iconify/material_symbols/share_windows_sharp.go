package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareWindowsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 15V7h10.15l-2.575-2.575L16 3l5 5l-5 5l-1.425-1.4L17.15 9H9v6H7Zm-4 6V4h2v15h12v-4h2v6H3Z"/>`),
		g.Group(children),
	)
}