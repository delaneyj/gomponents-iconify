package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitForkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5 4.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v6m0 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm10-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 0v1a3 3 0 0 1-3 3h-7"/>`),
		g.Group(children),
	)
}