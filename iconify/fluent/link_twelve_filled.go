package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.5 5.2a2.7 2.7 0 0 1 2.7-2.7h.05a.75.75 0 0 1 0 1.5H4.2A1.2 1.2 0 0 0 3 5.2v.6A1.2 1.2 0 0 0 4.2 7h.05a.75.75 0 0 1 0 1.5H4.2a2.7 2.7 0 0 1-2.7-2.7v-.6Zm9 0a2.7 2.7 0 0 0-2.7-2.7h-.05a.75.75 0 0 0 0 1.5h.05A1.2 1.2 0 0 1 9 5.2v.6A1.2 1.2 0 0 1 7.8 7h-.05a.75.75 0 0 0 0 1.5h.05a2.7 2.7 0 0 0 2.7-2.7v-.6Zm-5.75-.45a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Z"/>`),
		g.Group(children),
	)
}