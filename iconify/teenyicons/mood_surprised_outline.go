package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodSurprisedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 5.5h1m5 0h1m-3.5 9a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm-.5-7h1a1.5 1.5 0 1 1 0 3H7a1.5 1.5 0 1 1 0-3Z"/>`),
		g.Group(children),
	)
}