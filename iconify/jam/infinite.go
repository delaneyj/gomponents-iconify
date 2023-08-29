package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 9a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7zM11 5.5a3.5 3.5 0 1 0 .668-2.057a6.483 6.483 0 0 0-1.001-1.887A5.5 5.5 0 1 1 10 8.663A5.5 5.5 0 1 1 11 5.5z"/>`),
		g.Group(children),
	)
}