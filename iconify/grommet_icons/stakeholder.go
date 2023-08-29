package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stakeholder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m14 9l4.5 2L23 9V4l-4.5-2L14 4v5ZM7 7a4 4 0 1 0 0 8a4 4 0 0 0 0-8ZM1 23v-2c0-4 2.5-6 6-6s6 2 6 6v2H1ZM14 4l4.5 2L23 4m-4.5 2v5v-5Z"/>`),
		g.Group(children),
	)
}