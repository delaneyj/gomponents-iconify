package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 20l-7.5-7.428A5 5 0 1 1 12 6.006a5 5 0 1 1 7.893 6.139M19 22v-6m3 3l-3-3l-3 3"/>`),
		g.Group(children),
	)
}