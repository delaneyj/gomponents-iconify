package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapsuleF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 9.858H0v-4a5 5 0 1 1 10 0v4zm0 2v4a5 5 0 0 1-10 0v-4h10z"/>`),
		g.Group(children),
	)
}