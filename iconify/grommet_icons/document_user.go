package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 9V1H19.5L23 4.5V23h-7m2-22v5h5M8 11a3 3 0 1 0 0 6a3 3 0 0 0 0-6ZM3 23v-1c0-4 3-5 5-5s5 1 5 5v1H3Z"/>`),
		g.Group(children),
	)
}