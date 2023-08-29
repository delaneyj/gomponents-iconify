package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourKFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm8.5 10.5V12h-1V9H9v3H7.5V9H6v4.5h3V15h1.5v-1.5h1ZM18 15l-2.25-3L18 9h-1.75l-1.75 2.25V9H13v6h1.5v-2.25L16.25 15H18Z"/>`),
		g.Group(children),
	)
}