package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.007 11.042a4.75 4.75 0 0 1 4.432-3.04h23.807c3.336 0 5.632 3.348 4.431 6.46l-8.684 22.5A4.75 4.75 0 0 1 31.562 40H7.754c-3.335 0-5.632-3.348-4.43-6.46l8.683-22.5Z"/>`),
		g.Group(children),
	)
}