package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.727 18.242L4.792 27.208l8.966-8.966l-4.483-4.484l17.933-8.966l-8.966 8.966l4.485 4.484z"/>`),
		g.Group(children),
	)
}