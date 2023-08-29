package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteOffTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.293 4L2.146 2.854a.5.5 0 1 1 .708-.708l15 15a.5.5 0 0 1-.708.708l-1.998-2A3 3 0 0 1 12.272 18H7.728a3 3 0 0 1-2.98-2.656L3.554 5H2.5a.5.5 0 0 1 0-1h.793ZM12 12.707l-1-1V14a.5.5 0 0 0 1 0v-1.293Zm-3-3l-1-1V14a.5.5 0 0 0 1 0V9.707ZM12 8v1.879l3.481 3.48L16.446 5H17.5a.5.5 0 0 0 0-1h-5a2.5 2.5 0 0 0-5 0H6.121L11 8.879V8a.5.5 0 0 1 1 0Zm-.5-4h-3a1.5 1.5 0 1 1 3 0Z"/>`),
		g.Group(children),
	)
}