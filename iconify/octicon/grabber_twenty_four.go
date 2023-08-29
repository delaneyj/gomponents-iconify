package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrabberTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM9 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM9 18a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm6 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}