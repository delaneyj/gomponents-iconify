package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapRoadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMapRoadTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M42 4H6a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path d="m10 22l28-10"/><path stroke-dasharray="2 6" d="m10 29l28-10"/><path d="m10 36l28-10m-22-7l-3-8m24 25l-3-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMapRoadTwo0)"/>`),
		g.Group(children),
	)
}