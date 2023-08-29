package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M226.878 16v243.274a120.023 120.023 0 0 0 30 236.71a119.996 119.996 0 0 0 29.998-236.242v-93.747h59.998v-44.998h-59.998v-45h89.997V16H226.878zm25.312 299.99a60.01 60.01 0 0 1 2.343 0a59.998 59.998 0 0 1 2.343 0a59.998 59.998 0 0 1 0 119.996a60.044 60.044 0 0 1-4.688-119.996z"/>`),
		g.Group(children),
	)
}