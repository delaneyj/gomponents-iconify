package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23.923 0L10.959 1.893v9.588l12.964-.102V0ZM0 3.398l.009 8.155l9.773-.055l-.004-9.432L0 3.398Zm.008 17.283l9.773 1.344l-.008-9.44l-9.766-.063l.001 8.159Zm10.951 1.49L23.923 24l.004-11.326l-12.986-.022l.018 9.519Z"/>`),
		g.Group(children),
	)
}