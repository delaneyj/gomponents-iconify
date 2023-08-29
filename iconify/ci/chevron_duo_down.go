package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDuoDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18.424l-6.01-6.01l1.414-1.415l4.6 4.6l4.6-4.6l1.406 1.415l-6.009 6.01H12ZM12 13L5.99 6.989l1.414-1.414l4.6 4.6l4.6-4.6l1.406 1.414l-6.009 6.01L12 13Z"/>`),
		g.Group(children),
	)
}