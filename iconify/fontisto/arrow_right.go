package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.034 22.933L8 22.899l10.608-10.634L8 1.632l.034-.034h5.319L24 12.264L13.353 22.931z"/><path fill="currentColor" d="M0 10.488h18.666v3.555H0z"/>`),
		g.Group(children),
	)
}