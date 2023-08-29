package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordStopTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5Zm6 2A5 5 0 1 1 1 6a5 5 0 0 1 10 0Zm-1 0a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}