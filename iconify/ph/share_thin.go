package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.83 106.83l-48 48a4 4 0 0 1-5.66-5.66L214.34 108H165a92 92 0 0 0-89.11 69a4 4 0 0 1-3.89 3a3.87 3.87 0 0 1-1-.13a4 4 0 0 1-2.87-4.87A99.93 99.93 0 0 1 165 100h49.36l-41.19-41.17a4 4 0 0 1 5.66-5.66l48 48a4 4 0 0 1 0 5.66ZM192 212H40a4 4 0 0 1-4-4V88a4 4 0 0 0-8 0v120a12 12 0 0 0 12 12h152a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}