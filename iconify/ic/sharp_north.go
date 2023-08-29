package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpNorth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 9l1.41 1.41L11 5.83V22h2V5.83l4.59 4.59L19 9l-7-7l-7 7z"/>`),
		g.Group(children),
	)
}