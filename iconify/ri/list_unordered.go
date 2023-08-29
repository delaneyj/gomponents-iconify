package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUnordered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4h13v2H8V4ZM4.5 6.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm0 7a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm0 6.9a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM8 11h13v2H8v-2Zm0 7h13v2H8v-2Z"/>`),
		g.Group(children),
	)
}