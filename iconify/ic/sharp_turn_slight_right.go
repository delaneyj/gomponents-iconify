package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTurnSlightRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.34 6V4H18v5.66h-2V7.41l-5 5V20H9v-8.41L14.59 6z"/>`),
		g.Group(children),
	)
}