package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathematicsDisplayInline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 13H0V7h4zm12-6h4v6h-4zM6 6l3 4l-3 4h8v-3h-2v1H9.5l1.5-2l-1.5-2H12v1h2V6z"/>`),
		g.Group(children),
	)
}