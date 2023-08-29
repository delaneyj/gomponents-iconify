package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12c.001 5.515 4.487 10.001 10 10.001c5.514 0 10-4.486 10.001-10.001c0-5.514-4.486-10-10.001-10zm0 18.001c-4.41 0-7.999-3.589-8-8.001c0-4.411 3.589-8 8-8c4.412 0 8.001 3.589 8.001 8c-.001 4.412-3.59 8.001-8.001 8.001z"/>`),
		g.Group(children),
	)
}