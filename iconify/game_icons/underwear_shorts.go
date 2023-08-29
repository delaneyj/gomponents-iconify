package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnderwearShorts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m71.1 67.17l-4.28 38.43l379.28.7l-3.8-39.13zm-6.28 56.43L24.04 423.5c55.85 22.4 114.06 20.6 173.86 1L234 282.6c12.1 9 25.5 5.3 32.2-1l39.9 146.1c61.9 24.1 132 19.5 181.9 6l-40.2-309.4z"/>`),
		g.Group(children),
	)
}