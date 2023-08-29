package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 3a1 1 0 0 0 0 2v5.62l-4.789 5.387C1.491 17.942 2.865 21 5.454 21h13.092c2.589 0 3.962-3.058 2.242-4.993L16 10.62V5a1 1 0 1 0 0-2H8z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}