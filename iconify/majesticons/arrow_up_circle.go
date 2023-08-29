package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18zm.707-5.293a1 1 0 0 1-1.415-1.414L12.585 13H9a1 1 0 1 1 0-2h3.586l-1.293-1.293a1 1 0 0 1 1.415-1.414l3 3a1 1 0 0 1 0 1.414l-3 3z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}