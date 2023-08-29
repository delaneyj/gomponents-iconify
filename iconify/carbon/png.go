package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Png(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 23h-6a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h6v2h-6v10h4v-4h-2v-2h4zm-12-4L14.32 9H12v14h2V13l3.68 10H20V9h-2v10zM4 23H2V9h6a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H4zm0-7h4v-5H4z"/>`),
		g.Group(children),
	)
}