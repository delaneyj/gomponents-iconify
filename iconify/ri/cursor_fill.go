package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.91 12.36L17 20.854l-2.818 1.026l-3.092-8.493l-4.172 3.155l1.49-14.909l10.726 10.463l-5.225.264Z"/>`),
		g.Group(children),
	)
}