package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Zm-7.5-4.312c0-1.097-1.43-1.518-2.024-.596l-4.157 6.443A.95.95 0 0 0 9.117 15H13v1.25a.75.75 0 0 0 1.5 0V15h.75a.75.75 0 0 0 0-1.5h-.75V7.688ZM13 9.046V13.5h-2.874L13 9.046Z"/>`),
		g.Group(children),
	)
}