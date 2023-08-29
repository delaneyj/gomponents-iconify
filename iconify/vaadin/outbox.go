package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 5v6h4V5h2L8 0L4 5z"/><path fill="currentColor" d="M13 2h-2l.9 1h.4l2.6 8H11v2H5v-2H1.1l2.6-8h.4L5 2H3l-3 9v5h16v-5z"/>`),
		g.Group(children),
	)
}