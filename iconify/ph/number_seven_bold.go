package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSevenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m179.49 51.45l-48 160A12 12 0 0 1 120 220a11.82 11.82 0 0 1-3.45-.51a12 12 0 0 1-8-14.94L151.87 60H88a12 12 0 0 1 0-24h80a12 12 0 0 1 11.49 15.45Z"/>`),
		g.Group(children),
	)
}