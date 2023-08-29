package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 6.19l-2-2V2h10l-3.5 7H17l-1.82 3.37l-3.75-3.75L13.76 4H9v2.19m10 12.54L17.73 20l-4.32-4.32L10 22v-8H7V9.27l-5-5L3.27 3L19 18.73Z"/>`),
		g.Group(children),
	)
}