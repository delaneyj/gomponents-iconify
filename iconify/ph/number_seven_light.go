package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSevenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m173.75 49.72l-48 160a6 6 0 1 1-11.5-3.45L159.94 54H88a6 6 0 0 1 0-12h80a6 6 0 0 1 5.75 7.72Z"/>`),
		g.Group(children),
	)
}