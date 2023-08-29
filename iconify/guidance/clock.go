package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m17.5 17.5l-.404-.566A19.67 19.67 0 0 0 12 12l.055-.493A34.15 34.15 0 0 0 12 3.5M22.498 12c0-5.798-4.7-10.498-10.498-10.498c-5.798 0-10.498 4.7-10.498 10.498c0 5.798 4.7 10.498 10.498 10.498c5.798 0 10.498-4.7 10.498-10.498Z"/>`),
		g.Group(children),
	)
}