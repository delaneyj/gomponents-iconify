package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MickeyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18.501 2a4.5 4.5 0 0 1 .878 8.913a8 8 0 1 1-15.374 3.372L4 14l.005-.285a7.991 7.991 0 0 1 .615-2.803a4.5 4.5 0 0 1-3.187-6.348a4.505 4.505 0 0 1 3.596-2.539l.225-.018L5.535 2l.244.009a4.5 4.5 0 0 1 4.215 4.247a8.001 8.001 0 0 1 4.013 0A4.5 4.5 0 0 1 18.5 2z"/></g>`),
		g.Group(children),
	)
}