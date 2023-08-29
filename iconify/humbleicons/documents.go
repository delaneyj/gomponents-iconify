package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Documents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v7h7M6 7H5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-1M8 4v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V9.389a1 1 0 0 0-.263-.676l-4.94-5.389A1 1 0 0 0 14.06 3H9a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}