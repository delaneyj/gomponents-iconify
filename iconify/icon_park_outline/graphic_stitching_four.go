package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicStitchingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM9 44a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm5-40H4v10h10V4Zm30 30H34v10h10V34ZM34 9H14m20 30H14m-5-5V14m30 20V14"/>`),
		g.Group(children),
	)
}