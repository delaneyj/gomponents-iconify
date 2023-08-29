package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundLan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22h4c1.1 0 2-.9 2-2v-3c0-1.1-.9-2-2-2h-1v-2c0-1.1-.9-2-2-2h-3V9h1c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2h-4c-1.1 0-2 .9-2 2v3c0 1.1.9 2 2 2h1v2H8c-1.1 0-2 .9-2 2v2H5c-1.1 0-2 .9-2 2v3c0 1.1.9 2 2 2h4c1.1 0 2-.9 2-2v-3c0-1.1-.9-2-2-2H8v-2h8v2h-1c-1.1 0-2 .9-2 2v3c0 1.1.9 2 2 2z"/>`),
		g.Group(children),
	)
}