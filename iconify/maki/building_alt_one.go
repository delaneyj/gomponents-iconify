package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11 13.5v-9c0-.3-.2-.5-.5-.5H9V1L5 2.1v11.4H2v.5h11v-.5h-2zm-4 0V3h1v10.5H7z"/>`),
		g.Group(children),
	)
}