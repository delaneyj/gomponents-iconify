package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3l5.16 5.16L2 12h3v8h6v-6h1.11l.89.89V20h5.11l2.73 2.73l1.27-1.27M9 12v6H7v-7.81l.68-.62L10.11 12H9m6 6v-1.11L16.11 18H15M10.36 7.16L8.95 5.75L12 3l10 9h-3v3.8l-2-2v-3.61l-5-4.5l-1.64 1.47Z"/>`),
		g.Group(children),
	)
}