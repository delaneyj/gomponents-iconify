package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSignBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRoadSignBoth0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 8v8h28l4-4l-4-4H10Zm28 15v8H10l-4-4l4-4h28Z"/><path stroke-linecap="round" d="M24 31v13m0-28v7m0-19v4m-5 36h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRoadSignBoth0)"/>`),
		g.Group(children),
	)
}