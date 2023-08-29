package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M10 2a3 3 0 0 1 3 3v1h1a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1V5a3 3 0 0 1 3-3zm0 8.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM10 4a1 1 0 0 0-1 1v1h2V5a1 1 0 0 0-1-1z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}