package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSpellcheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.45 16h2.09L9.43 3H7.57L2.46 16h2.09l1.12-3h5.64l1.14 3zm-6.02-5L8.5 5.48L10.57 11H6.43zm15.16.59l-8.09 8.09L9.83 16l-1.41 1.41l5.09 5.09L23 13l-1.41-1.41z"/>`),
		g.Group(children),
	)
}