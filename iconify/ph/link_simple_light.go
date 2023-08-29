package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164.25 91.75a6 6 0 0 1 0 8.49l-64 64a6 6 0 0 1-8.49-8.48l64-64a6 6 0 0 1 8.49-.01ZM214.2 41.8a54.07 54.07 0 0 0-76.38 0l-30.07 30.05a6 6 0 0 0 8.49 8.49l30.07-30.06a42 42 0 0 1 59.41 59.41l-30.08 30.07a6 6 0 1 0 8.49 8.49l30.07-30.07a54 54 0 0 0 0-76.38Zm-74.44 133.84l-30.07 30.08a42 42 0 0 1-59.41-59.41l30.06-30.07a6 6 0 0 0-8.49-8.49l-30 30.07a54 54 0 0 0 76.38 76.39l30.07-30.08a6 6 0 0 0-8.49-8.49Z"/>`),
		g.Group(children),
	)
}