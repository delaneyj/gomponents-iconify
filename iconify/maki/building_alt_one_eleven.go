package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingAltOneEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8 9.5v-5c0-.3-.2-.5-.5-.5H6V1L3 2v7.5H2v.5h7v-.5H8zm-3 0H4V3h1v6.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}