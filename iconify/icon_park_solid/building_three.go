package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBuildingThree0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="m24 8l20 13v23H4V21L24 8Z" clip-rule="evenodd"/><path stroke="#000" d="M20 44V23l-8 5v16m16 0V23l8 5v16"/><path stroke="#fff" d="M41 44H8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBuildingThree0)"/>`),
		g.Group(children),
	)
}