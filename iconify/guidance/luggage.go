package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5.5 4.5v17m13-17v17m-2.027-17H23.5v.25l-.055.31a45.687 45.687 0 0 0 0 15.88l.055.31v.25H.5v-.25l.055-.31a45.682 45.682 0 0 0 0-15.88L.5 4.75V4.5h7.027m8.945 0a4.5 4.5 0 0 0-8.945 0m8.945 0H7.527"/>`),
		g.Group(children),
	)
}