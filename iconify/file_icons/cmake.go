package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cmake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M272.728 22.068L512 497.945L306.578 413.79l-33.85-391.723zm-266.892 446l256.677-218.594L240.924 0L5.836 468.069zm146.217-85.645L0 512h467.856L152.053 382.423z"/>`),
		g.Group(children),
	)
}