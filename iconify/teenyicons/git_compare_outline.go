package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCompareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m8.5.5l-2 2m0 0l2 2m-2-2h3a3 3 0 0 1 3 3v5m-10-6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 0v5a3 3 0 0 0 3 3H8m-1.5 2l2-2l-2-2m6 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}