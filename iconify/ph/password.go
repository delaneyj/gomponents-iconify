package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Password(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M48 56v144a8 8 0 0 1-16 0V56a8 8 0 0 1 16 0Zm84 54.5l-20 6.5V96a8 8 0 0 0-16 0v21l-20-6.5a8 8 0 0 0-5 15.22l20 6.49l-12.34 17a8 8 0 1 0 12.94 9.4l12.34-17l12.34 17a8 8 0 1 0 12.94-9.4l-12.34-17l20-6.49A8 8 0 0 0 132 110.5Zm106 5.14a8 8 0 0 0-10-5.14l-20 6.5V96a8 8 0 0 0-16 0v21l-20-6.49a8 8 0 0 0-4.95 15.22l20 6.49l-12.34 17a8 8 0 1 0 12.94 9.4l12.34-17l12.34 17a8 8 0 1 0 12.94-9.4l-12.34-17l20-6.49a8 8 0 0 0 5.07-10.09Z"/>`),
		g.Group(children),
	)
}