package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M270.877 444.542C576.857 496.618 318.44 29.007 23.097 25.68C447.57-7.506 696.864 640.745 270.878 444.54z"/>`),
		g.Group(children),
	)
}