package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm-.707-12.707a1 1 0 0 1 1.414 1.415L11.414 11H15a1 1 0 1 1 0 2h-3.585l1.292 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.415l3-3z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}