package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M11 7v8h4l4 5m-8-9h5m-9 .5a5 5 0 1 0 6 7.5"/></g>`),
		g.Group(children),
	)
}