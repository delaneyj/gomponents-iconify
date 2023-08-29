package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.139 2.528a1.492 1.492 0 0 0-2.278 0l-6.62 7.803c-.553.651-.093 1.654.759 1.654h3.01v5.012A1 1 0 0 0 7.006 18h5.986a1 1 0 0 0 .998-1.003v-5.012H17c.85 0 1.31-1.003.759-1.654l-6.621-7.803Z"/>`),
		g.Group(children),
	)
}