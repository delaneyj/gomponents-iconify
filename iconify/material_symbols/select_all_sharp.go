package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectAllSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17V7h10v10H7Zm2-2h6V9H9v6Zm-2 6v-2h2v2H7ZM7 5V3h2v2H7Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2ZM3 5V3h2v2H3Zm2 16H3v-2h2v2Zm14 0v-2h2v2h-2Zm2-16h-2V3h2v2ZM3 17v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm16 8v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Z"/>`),
		g.Group(children),
	)
}