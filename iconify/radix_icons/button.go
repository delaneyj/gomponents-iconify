package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Button(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5h11a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1ZM0 6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V6Zm4.5.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm2.25.75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0Zm3.75-.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}