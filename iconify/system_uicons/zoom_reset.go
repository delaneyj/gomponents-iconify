package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomReset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 8.5a5 5 0 1 0 1.057-3.074"/><path d="M4.5 1.5v4h4m9 12l-5.379-5.379"/></g>`),
		g.Group(children),
	)
}