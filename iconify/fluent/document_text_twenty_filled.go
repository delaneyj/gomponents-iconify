package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentTextTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 6.5V2H5.5A1.5 1.5 0 0 0 4 3.5v13A1.5 1.5 0 0 0 5.5 18h9a1.5 1.5 0 0 0 1.5-1.5V8h-4.5A1.5 1.5 0 0 1 10 6.5ZM6.5 10h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1Zm0 2h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1Zm0 2h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1 0-1ZM11 6.5V2.25L15.75 7H11.5a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}