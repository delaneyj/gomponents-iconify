package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSimpleOpenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m227.33 91l-96-64a6 6 0 0 0-6.66 0l-96 64A6 6 0 0 0 26 96v104a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V96a6 6 0 0 0-2.67-5ZM128 39.21l85.43 57l-69.9 49.79h-31.06l-69.9-49.83ZM216 202H40a2 2 0 0 1-2-2v-92.35l69.06 49.24a6.06 6.06 0 0 0 3.49 1.11h34.9a6.06 6.06 0 0 0 3.49-1.11L218 107.65V200a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}