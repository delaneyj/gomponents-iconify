package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarSmileFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.999.5l4.226 6.183l7.186 2.109l-4.575 5.93l.216 7.486l-7.053-2.518l-7.054 2.518l.216-7.486l-4.575-5.93l7.187-2.109L11.999.5Zm-2 11.5h-2a4 4 0 0 0 7.995.2l.005-.2h-2a2 2 0 0 1-3.995.15L10 12Z"/>`),
		g.Group(children),
	)
}