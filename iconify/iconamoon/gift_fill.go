package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 1a3.5 3.5 0 0 0-3.163 5H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1v7a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-2.337A3.5 3.5 0 0 0 15.5 1A4.491 4.491 0 0 0 12 2.671A4.491 4.491 0 0 0 8.5 1ZM13 20h5a1 1 0 0 0 1-1v-7h-6v8Zm-2-8v8H6a1 1 0 0 1-1-1v-7h6Zm4.5-6a1.5 1.5 0 0 0 0-3A2.5 2.5 0 0 0 13 5.5V6h2.5ZM11 6v-.5A2.5 2.5 0 0 0 8.5 3a1.5 1.5 0 1 0 0 3H11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}