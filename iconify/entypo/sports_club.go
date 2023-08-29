package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsClub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6 13.5l4 2.5l4-2.5V5H6v8.5zM4.5 10a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm13-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM4.485 6.199A6.71 6.71 0 0 1 10 3.3a6.715 6.715 0 0 1 5.456 2.823a1.4 1.4 0 0 0 2.281-1.624A9.518 9.518 0 0 0 10 .5a9.506 9.506 0 0 0-7.817 4.107a1.402 1.402 0 0 0 .355 1.948a1.401 1.401 0 0 0 1.947-.356zm10.971 7.678A6.713 6.713 0 0 1 10 16.7a6.71 6.71 0 0 1-5.515-2.899a1.4 1.4 0 0 0-2.302 1.592A9.506 9.506 0 0 0 10 19.5a9.518 9.518 0 0 0 7.737-3.999a1.401 1.401 0 0 0-2.281-1.624z"/>`),
		g.Group(children),
	)
}