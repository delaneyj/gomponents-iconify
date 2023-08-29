package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4h10a2 2 0 0 1 2 2v10m-.586 3.414A2 2 0 0 1 18 20H6a2 2 0 0 1-2-2V6c0-.547.22-1.043.576-1.405"/><path d="m7 14l3-3l2 2l.5-.5m2-2l.5-.5l2 2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}