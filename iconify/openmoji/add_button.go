package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36.073" cy="35.952" r="22.77" fill="#B1CC33"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M36.073 22.327v27.25m13.625-13.625h-27.25"/><circle cx="36.073" cy="35.952" r="23"/></g>`),
		g.Group(children),
	)
}