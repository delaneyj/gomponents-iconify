package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5.27L3.28 4L20 20.72L18.73 22l-6-6H9v-2l.79-.94L2 5.27M23 12h-6v-2l3.39-4H17V4h6v2l-3.38 4H23v2M9.82 8H15v2l-1.46 1.72L9.82 8M7 20H1v-2l3.39-4H1v-2h6v2l-3.38 4H7v2Z"/>`),
		g.Group(children),
	)
}