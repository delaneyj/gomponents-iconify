package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maxcdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m5 6l2 5l-3 15h5l3-15h4l-3 15h5l3-15h4l-3 15h5l2.75-13.742C30.395 9.02 27.922 6 24.617 6z"/>`),
		g.Group(children),
	)
}