package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20V4H4v15a1 1 0 0 0 1 1h11Zm3 2H5a3 3 0 0 1-3-3V3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v7h4v9a3 3 0 0 1-3 3Zm-1-10v7a1 1 0 1 0 2 0v-7h-2ZM6 6h6v6H6V6Zm2 2v2h2V8H8Zm-2 5h8v2H6v-2Zm0 3h8v2H6v-2Z"/>`),
		g.Group(children),
	)
}