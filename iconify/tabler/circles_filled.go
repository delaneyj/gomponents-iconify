package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M6.5 12a5 5 0 1 1-4.995 5.217L1.5 17l.005-.217A5 5 0 0 1 6.5 12zm11 0a5 5 0 1 1-4.995 5.217L12.5 17l.005-.217A5 5 0 0 1 17.5 12zM12 2a5 5 0 1 1-4.995 5.217L7 7l.005-.217A5 5 0 0 1 12 2z"/></g>`),
		g.Group(children),
	)
}