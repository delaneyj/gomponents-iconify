package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 5l-3-3m0 0L9 5m3-3v12M6 9H4v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9h-2"/>`),
		g.Group(children),
	)
}