package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarLockTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 3A2.5 2.5 0 0 1 17 5.5v3.901a2.98 2.98 0 0 0-1-.36V7H4v7.5A1.5 1.5 0 0 0 5.5 16H11v1H5.5A2.5 2.5 0 0 1 3 14.5v-9A2.5 2.5 0 0 1 5.5 3h9Zm0 1h-9A1.5 1.5 0 0 0 4 5.5V6h12v-.5A1.5 1.5 0 0 0 14.5 4Zm-1 8v1H13a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-.5v-1a2 2 0 1 0-4 0Zm1 1v-1a1 1 0 1 1 2 0v1h-2Zm1 2.25a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}