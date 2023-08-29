package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuciferCannon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M376 76c30 60-120 30-210 75c150 0 270 15 270 105S316 361 166 361c90 45 240 15 210 75c90-15 120-120 120-180S466 91 376 76zm-95.625 105.938C216.005 182.577 127.562 203.5 16 256c255 120 390 75 390 0c0-42.188-42.865-74.886-125.625-74.063z"/>`),
		g.Group(children),
	)
}