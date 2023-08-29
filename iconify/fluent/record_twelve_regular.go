package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm5-3A5 5 0 1 1 1 6a5 5 0 0 1 10 0Zm-1 0a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}