package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBuildingFour0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" d="m17 14l27 10v20H17V14Z" clip-rule="evenodd"/><path stroke="#fff" d="M17 14L4 24v20h13"/><path stroke="#000" d="M35 44V32l-9-3v15"/><path stroke="#fff" d="M44 44H17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBuildingFour0)"/>`),
		g.Group(children),
	)
}