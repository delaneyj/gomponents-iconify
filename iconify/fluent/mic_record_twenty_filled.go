package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicRecordTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 10a.5.5 0 0 0-1 0a5.5 5.5 0 0 0 5 5.478V17.5a.5.5 0 0 0 1 0v-.706A5.48 5.48 0 0 1 9 14.5A4.5 4.5 0 0 1 4.5 10ZM12 5v4.6a5.514 5.514 0 0 0-2.79 3.393A3 3 0 0 1 6 10V5a3 3 0 0 1 6 0Zm5 9.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm2 0a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-8 0a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Z"/>`),
		g.Group(children),
	)
}