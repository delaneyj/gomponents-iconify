package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10ZM10.75 7a.75.75 0 0 0-.748.697l-.25 3.527a.75.75 0 0 0 .791.801h.015l.047-.003a25.755 25.755 0 0 1 .743-.028c.462-.011.936-.008 1.184.026a2 2 0 1 1-2.097 2.816a.75.75 0 1 0-1.362.628a3.5 3.5 0 1 0 3.667-4.93h-.003c-.385-.053-.978-.05-1.425-.04h-.005L11.45 8.5h2.802a.75.75 0 0 0 0-1.5H10.75Z"/>`),
		g.Group(children),
	)
}