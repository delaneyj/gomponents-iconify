package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightClub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M47 9H32s-1 8 1 16c.804 1.34 2.79 2.062 5 2.245V37h-.333L34 40h11l-3.584-3H41v-9.828c2.12-.275 4.063-1.014 5-2.172c2-7 1-16 1-16zm-1 6H33v-5h13v5zM21.5 37H16v-8.5L27 12H1l11 16.5V37H6.5a1.5 1.5 0 1 0 0 3h15a1.5 1.5 0 1 0 0-3z"/>`),
		g.Group(children),
	)
}