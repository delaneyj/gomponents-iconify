package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathematicsDisplayDefault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 5H4l3 5l-3 5h10v-3h-2v1H8.2l1.8-3l-1.8-3H12v1h2V5zM1 9h3v2H1zm15 0h3v2h-3z"/>`),
		g.Group(children),
	)
}