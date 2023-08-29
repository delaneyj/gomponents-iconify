package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 5l-1.414-1.414A2 2 0 0 0 11.172 3H5a2 2 0 0 0-2 2v6.172a2 2 0 0 0 .586 1.414L5 14m14-4l1.586 1.586a2 2 0 0 1 0 2.828l-6.172 6.172a2 2 0 0 1-2.828 0L10 19M7 7h.001M21 3L3 21"/>`),
		g.Group(children),
	)
}