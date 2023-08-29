package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="m3.75 13.2l6.5-5.5l4 3"/><circle cx="5.25" cy="6.25" r=".5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}