package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFastfood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 5V1h-2v4h-5l.23 2.31C14.9 8.16 18 10.77 18 15l.02 8h3.18L23 5h-5zM1 21h15v2H1zM8.5 8.99C4.75 8.99 1 11 1 15h15c0-4-3.75-6.01-7.5-6.01zM1 17h15v2H1z"/>`),
		g.Group(children),
	)
}