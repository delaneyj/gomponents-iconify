package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.502 4.502 0 0 0 6.5 19h12c.159 0 .316-.01.469-.031m-8.203-13.86A6.992 6.992 0 0 1 16.95 7.05A6.978 6.978 0 0 1 19 12.035a3.5 3.5 0 0 1 2.917 4.225"/><path d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}