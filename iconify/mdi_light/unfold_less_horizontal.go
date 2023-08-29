package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLessHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.74 5.29L11.5 9.54L7.26 5.29l.7-.7l3.54 3.53l3.54-3.53l.7.7m0 14.42l-.7.7l-3.54-3.53l-3.54 3.53l-.7-.7l4.24-4.25l4.24 4.25Z"/>`),
		g.Group(children),
	)
}