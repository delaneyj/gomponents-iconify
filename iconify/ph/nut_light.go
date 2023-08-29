package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 82a46 46 0 1 0 46 46a46.06 46.06 0 0 0-46-46Zm0 80a34 34 0 1 1 34-34a34 34 0 0 1-34 34Zm94.72-94.09l-88-48.18a13.9 13.9 0 0 0-13.44 0l-88 48.18A14 14 0 0 0 26 80.18v95.64a14 14 0 0 0 7.28 12.27l88 48.18a13.92 13.92 0 0 0 13.44 0l88-48.18a14 14 0 0 0 7.28-12.27V80.18a14 14 0 0 0-7.28-12.27ZM218 175.82a2 2 0 0 1-1 1.75l-88 48.18a2 2 0 0 1-1.92 0L39 177.57a2 2 0 0 1-1-1.75V80.18a2 2 0 0 1 1-1.75l88-48.18a2 2 0 0 1 1.92 0l88 48.18a2 2 0 0 1 1 1.75Z"/>`),
		g.Group(children),
	)
}