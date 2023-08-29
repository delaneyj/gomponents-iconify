package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataHistogramSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 4v9h3V4a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1Zm-1 3V4a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1H12a2 2 0 0 1 2 2v6.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V9a2 2 0 0 1 2-2h1.5Zm0 6V8H4a1 1 0 0 0-1 1v4h2.5Zm5 0H13V7a1 1 0 0 0-1-1h-1.5v7Z"/>`),
		g.Group(children),
	)
}