package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Future(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 0v2C6.955 2 2 6.955 2 13s4.955 11 11 11s11-4.955 11-11c0-2.947-1.11-5.61-3.094-7.594L19.5 6.813C21.117 8.428 22 10.546 22 13c0 4.955-4.045 9-9 9s-9-4.045-9-9s4.045-9 9-9v2l5-3l-5-3zm2.094 6.563l-2.5 5A1.483 1.483 0 0 0 11.5 13a1.5 1.5 0 0 0 1.5 1.5h.063l3.218 3.219l1.438-1.438l-3.219-3.218V13c0-.197-.056-.39-.125-.563l2.531-5l-1.812-.875z"/>`),
		g.Group(children),
	)
}