package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelPlacesMountainSunLandmarkMountainPlace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m13.5 9.26l-3.79-5a1 1 0 0 0-1.42 0L.5 13.5h13"/><path d="m4.69 8.5l1.47 1a1 1 0 0 0 1.16 0l.89-.67a1 1 0 0 1 1.2 0l.89.67a1 1 0 0 0 1.15 0l1.48-1"/><circle cx="3" cy="3" r="2.5"/></g>`),
		g.Group(children),
	)
}