package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.998 7l2.31-4h.7c.548 0 .992.445.992.993v16.014a1 1 0 0 1-.992.993H2.992A.993.993 0 0 1 2 20.007V3.993A1 1 0 0 1 2.992 3h3.006l-2.31 4h2.31l2.31-4h3.69l-2.31 4h2.31l2.31-4h3.69l-2.31 4h2.31Z"/>`),
		g.Group(children),
	)
}