package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBuildingFour0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" d="m17 14l27 10v20H17V14Z" clip-rule="evenodd"/><path d="M17 14L4 24v20h13m18 0V32l-9-3v15m18 0H17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBuildingFour0)"/>`),
		g.Group(children),
	)
}