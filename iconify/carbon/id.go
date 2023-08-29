package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Id(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M10 9h2v2h-2z" fill="currentColor"/><path d="M18 23h-4V9h4a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4zm-2-2h2a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-2z" fill="currentColor"/><path d="M10 13h2v10h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}