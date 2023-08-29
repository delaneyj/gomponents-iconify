package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm-4.25 9.25a.75.75 0 0 0-.102 1.493l.102.007h8.5a.75.75 0 0 0 .102-1.493l-.102-.007h-8.5Z"/>`),
		g.Group(children),
	)
}