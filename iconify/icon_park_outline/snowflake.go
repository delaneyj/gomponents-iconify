package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 4v40M6.725 14l34.64 20M6.72 33.977l34.56-19.954M12 10l3 9l-9 2m0 6l9 2l-3 9m24-28l-3 9l9 2m0 6l-9 2l3 9M18 7l6 6l6-6M18 41l6-6l6 6"/>`),
		g.Group(children),
	)
}