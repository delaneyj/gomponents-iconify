package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaMulti0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaMulti1" cx="56" cy="56" r="40"/><circle id="galaMulti2" cx="200" cy="56" r="40"/><circle id="galaMulti3" cx="56" cy="200" r="40"/><circle id="galaMulti4" cx="200" cy="200" r="40"/></g>`),
		g.Group(children),
	)
}