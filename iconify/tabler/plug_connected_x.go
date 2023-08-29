package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugConnectedX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 16l-4 4m-9-8l5 5l-1.5 1.5a3.536 3.536 0 1 1-5-5L7 12zm10 0l-5-5l1.5-1.5a3.536 3.536 0 1 1 5 5L17 12zM3 21l2.5-2.5m13-13L21 3m-11 8l-2 2m5 1l-2 2m5 0l4 4"/>`),
		g.Group(children),
	)
}