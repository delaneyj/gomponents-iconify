package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 10h-4V6h-2v4h-4v2h4v4h2v-4h4v-2z"/><path fill="currentColor" d="M28 10h-2V4a2.002 2.002 0 0 0-2-2H8a2.002 2.002 0 0 0-2 2v6H4a2.002 2.002 0 0 0-2 2v18h28V12a2.002 2.002 0 0 0-2-2ZM14 28v-6h4v6Zm6 0v-7a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v7H4V12h4V4h16v8h4v16Z"/>`),
		g.Group(children),
	)
}