package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlphaUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m75.313 130.313l51.883-51.881V496h32V78.432l51.883 51.882l22.627-22.627l-90.51-90.51l-90.509 90.51l22.626 22.626zM440 280H304v32h101.04L296 412.732V456h144v-32H330.96L440 323.268V280zM395.532 48h-47.064L289.8 224h33.73l18.67-56h59.6l18.667 56H454.2Zm-42.667 88l18.667-56h.936l18.667 56Z"/>`),
		g.Group(children),
	)
}