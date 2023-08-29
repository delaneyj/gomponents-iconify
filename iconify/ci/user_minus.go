package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19H2a6 6 0 0 1 12 0h-2a4 4 0 0 0-8 0Zm18-6h-8v-2h8v2ZM8 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-6a2 2 0 1 0 2 2.09v.4V8a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}