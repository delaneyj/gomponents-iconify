package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 11H16V6.5a.5.5 0 0 0-.5-.5H3V2.5a.5.5 0 0 0-1 0v19a.5.5 0 1 0 1 0V18h18.5a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5zM3 7h12v4H3V7zm18 10H3v-5h18v5z"/>`),
		g.Group(children),
	)
}