package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListIndefinite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4h8v2h-8V4Zm0 7h8v2h-8v-2Zm0 7h8v2h-8v-2Zm-6.5 1a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 2a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM5 6v3h3V6H5ZM3 4h7v7H3V4Z"/>`),
		g.Group(children),
	)
}