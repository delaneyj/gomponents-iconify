package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M42.516 55L25.54 45.889l-14.286-7.668c-1.672-.897-1.672-3.545 0-4.442L25.54 26.11L42.516 17"/><path d="m66 55l-17.055-9.111l-14.353-7.668c-1.68-.897-1.68-3.545 0-4.442l14.353-7.668L66 17M6 55V17"/></g>`),
		g.Group(children),
	)
}