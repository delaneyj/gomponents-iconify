package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Install(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M19 13.5v4L12 22l-7-4.5v-4m7 8.5v-8.5m6.5-5l-6.5-4L15.5 2L22 6l-3.5 2.5Zm-13 0l6.5-4L8.5 2L2 6l3.5 2.5Zm13 .5L12 13l3.5 2.5l6.5-4L18.5 9Zm-13 0l6.5 4l-3.5 2.5l-6.5-4L5.5 9Z"/>`),
		g.Group(children),
	)
}