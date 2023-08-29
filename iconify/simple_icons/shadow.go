package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shadow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12a11.97 11.97 0 0 0 3.918 8.87a4.457 4.457 0 0 1-.2-1.324a4.453 4.453 0 1 1 5.891 4.216c.773.156 1.572.238 2.391.238c6.627 0 12-5.373 12-12S18.627 0 12 0Z"/>`),
		g.Group(children),
	)
}