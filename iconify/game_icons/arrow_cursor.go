package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M123.193 29.635L121 406.18l84.31-82.836l65.87 159.02l67.5-27.96l-65.87-159.02L391 294.342z"/>`),
		g.Group(children),
	)
}