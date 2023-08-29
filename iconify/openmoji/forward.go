package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M46.815 33.428L55 25.214L46.815 17"/><path d="M32.329 54.498c-8.087 0-14.642-6.556-14.642-14.642s6.555-14.642 14.642-14.642h21.73"/></g>`),
		g.Group(children),
	)
}