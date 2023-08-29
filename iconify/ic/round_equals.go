package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundEquals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9.998H6a1 1 0 1 1 0-2h12a1 1 0 0 1 0 2zm0 6H6a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}