package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewAgenda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 213q9 0 15 6.5t6 15.5v128q0 8-6 14.5t-15 6.5H21q-8 0-14.5-6.5T0 363V235q0-9 6.5-15.5T21 213h363zm0-213q9 0 15 6.5t6 14.5v128q0 9-6 15.5t-15 6.5H21q-8 0-14.5-6.5T0 149V21q0-8 6.5-14.5T21 0h363z"/>`),
		g.Group(children),
	)
}