package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHi0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 43H17.476a.257.257 0 0 1-.218-.121L7.86 27.727a4.095 4.095 0 1 1 7.011-4.23l2.462 4.194V8.576a3 3 0 0 1 3.522-2.955L37.52 8.563A3 3 0 0 1 40 11.517V40a3 3 0 0 1-3 3Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHi0)"/>`),
		g.Group(children),
	)
}