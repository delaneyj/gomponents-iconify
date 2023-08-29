package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.5 9L7 14v20l8.5 5l8.5 5l8.5-5l8.5-5V14l-8.5-5L24 4l-8.5 5ZM41 14L24 24M7 14l17 10m0 20V24m8-5v20m9-15L24 34m0 0L7 24m9 15V19M32 9L16 19m16 0L16 9"/>`),
		g.Group(children),
	)
}