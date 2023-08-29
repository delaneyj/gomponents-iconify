package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 20v-8h-8v3h-5V7a2 2 0 0 0-2-2h-5V2H2v8h8V7h5v18a2 2 0 0 0 2 2h5v3h8v-8h-8v3h-5v-8h5v3ZM8 8H4V4h4Zm16 16h4v4h-4Zm0-10h4v4h-4Z"/>`),
		g.Group(children),
	)
}