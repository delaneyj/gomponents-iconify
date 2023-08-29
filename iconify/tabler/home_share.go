package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 21v-6a2 2 0 0 1 2-2h2c.247 0 .484.045.702.127"/><path d="M19 12h2l-9-9l-9 9h2v7a2 2 0 0 0 2 2h5m4 1l5-5m0 4.5V17h-4.5"/></g>`),
		g.Group(children),
	)
}