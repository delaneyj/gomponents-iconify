package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchOneFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 2a6.007 6.007 0 0 0-6 6h2a4 4 0 0 1 8 0h2a6.007 6.007 0 0 0-6-6Z"/><path fill="currentColor" d="M21 30h-4.44a4 4 0 0 1-2.708-1.057l-9.2-8.466a2.002 2.002 0 0 1 .118-3.055a2.074 2.074 0 0 1 2.658.173L11 20.857V8a2 2 0 0 1 4 0v7a2 2 0 0 1 4 0v1a2 2 0 0 1 4 0v1a2 2 0 0 1 4 0v7a6 6 0 0 1-6 6Z"/>`),
		g.Group(children),
	)
}