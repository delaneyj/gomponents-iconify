package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 2v2h2v6a4 4 0 0 1-8 0V4h2V2h-4v8a6.005 6.005 0 0 0 5 5.91V22a6 6 0 0 1-12 0v-6.142a4 4 0 1 0-2 0V22a8 8 0 0 0 16 0v-6.09A6.005 6.005 0 0 0 28 10V2ZM6 12a2 2 0 1 1 2 2a2.002 2.002 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}