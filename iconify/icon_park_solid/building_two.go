package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBuildingTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m21 13l-10 7v24"/><path fill="#fff" fill-rule="evenodd" d="m21 4l10 7v13l7 5v15H21V4Z" clip-rule="evenodd"/><path d="M4 44h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBuildingTwo0)"/>`),
		g.Group(children),
	)
}