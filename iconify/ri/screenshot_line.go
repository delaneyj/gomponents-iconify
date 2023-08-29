package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.993 14.407l-1.552 1.552a4 4 0 1 1-1.418-1.41l1.555-1.556l-4.185-4.185l1.415-1.415l4.185 4.185l4.189-4.189l1.414 1.414l-4.19 4.19l1.562 1.56a4 4 0 1 1-1.414 1.414l-1.561-1.56ZM7 20a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm2-7V5H5v8H3V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v9h-2Z"/>`),
		g.Group(children),
	)
}