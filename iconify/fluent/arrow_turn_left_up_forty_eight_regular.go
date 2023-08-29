package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftUpFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M40.75 35.5a1.25 1.25 0 1 1 0 2.5h-18.5A7.25 7.25 0 0 1 15 30.75V12.582l-6.854 7.04a1.25 1.25 0 0 1-1.792-1.744l9.25-9.5a1.25 1.25 0 0 1 1.792 0l9.25 9.5a1.25 1.25 0 0 1-1.792 1.744L17.5 12.069V30.75a4.75 4.75 0 0 0 4.75 4.75h18.5Z"/>`),
		g.Group(children),
	)
}