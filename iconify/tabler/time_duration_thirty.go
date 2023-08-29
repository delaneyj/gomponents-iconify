package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeDurationThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 10.5v3a1.5 1.5 0 0 0 3 0v-3a1.5 1.5 0 0 0-3 0zM8 9h1.5a1.5 1.5 0 0 1 0 3H9h.5a1.5 1.5 0 0 1 0 3H8m-5-3v.01M7.5 4.2v.01m0 15.59v.01M4.2 16.5v.01m0-9.01v.01"/><path d="M12 21a9 9 0 0 0 0-18"/></g>`),
		g.Group(children),
	)
}