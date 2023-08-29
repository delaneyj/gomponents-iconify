package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m0 47.526l503.929 503.906v493.521H318.627v107.521h562.769v-107.521H696.094v-493.52L1200 47.526H0z"/>`),
		g.Group(children),
	)
}