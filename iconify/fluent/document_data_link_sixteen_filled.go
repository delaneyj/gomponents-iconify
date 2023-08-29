package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataLinkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 5h3.25L10 1.25V4.5a.5.5 0 0 0 .5.5Zm0 1A1.5 1.5 0 0 1 9 4.5V1H5.5A1.5 1.5 0 0 0 4 2.5V9h2V6.5a.5.5 0 0 1 1 0V9h.5a3.5 3.5 0 0 1 2.45 6h2.55a1.5 1.5 0 0 0 1.5-1.5V6h-3.5Zm.5 6.5v-4a.5.5 0 0 1 1 0v4a.5.5 0 0 1-1 0ZM3.5 10a2.5 2.5 0 0 0 0 5H4a.5.5 0 0 0 0-1h-.5a1.5 1.5 0 0 1 0-3H4a.5.5 0 0 0 0-1h-.5ZM7 10a.5.5 0 0 0 0 1h.5a1.5 1.5 0 0 1 0 3H7a.5.5 0 0 0 0 1h.5a2.5 2.5 0 0 0 0-5H7Zm-4 2.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}