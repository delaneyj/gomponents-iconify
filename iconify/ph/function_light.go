package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 40a6 6 0 0 1-6 6h-29.29a26 26 0 0 0-25.58 21.35L135.19 122H184a6 6 0 0 1 0 12h-51l-10.33 56.8A38 38 0 0 1 85.29 222H56a6 6 0 0 1 0-12h29.29a26 26 0 0 0 25.58-21.35l9.94-54.65H72a6 6 0 0 1 0-12h51l10.33-56.8A38 38 0 0 1 170.71 34H200a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}