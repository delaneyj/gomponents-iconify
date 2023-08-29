package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm7-10a6.99 6.99 0 0 0-2.8-5.6L13 6a4.99 4.99 0 0 1 2 4a4.992 4.992 0 0 1-2 4l.58.87l.53.796A6.99 6.99 0 0 0 17 10zm-7 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/>`),
		g.Group(children),
	)
}