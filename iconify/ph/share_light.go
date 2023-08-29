package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m228.24 108.24l-48 48a6 6 0 0 1-8.48-8.48L209.51 110H165a89.94 89.94 0 0 0-87.17 67.5a6 6 0 0 1-11.62-3A101.94 101.94 0 0 1 165 98h44.53l-37.77-37.76a6 6 0 0 1 8.48-8.48l48 48a6 6 0 0 1 0 8.48ZM192 210H40a2 2 0 0 1-2-2V88a6 6 0 0 0-12 0v120a14 14 0 0 0 14 14h152a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}