package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnFlagLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.14 8.48L17 6L5.58 1.92zM1.22 0L0 1.22l3 3V19h2v-6.87l3.91-2L18.78 20L20 18.78z"/>`),
		g.Group(children),
	)
}