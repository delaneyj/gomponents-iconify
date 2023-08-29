package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Zm1.5-10a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0ZM12 5.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-2.5 4a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}