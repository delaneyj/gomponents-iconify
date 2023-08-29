package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="50" cy="22" r="5"/><circle cx="22" cy="38" r="5"/><circle cx="50" cy="50" r="5"/><path d="m27 40l18 8m0-23L27 35"/></g>`),
		g.Group(children),
	)
}