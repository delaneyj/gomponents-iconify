package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="2" y="21" rx=".6" transform="rotate(-90 2 21)"/><rect width="7" height="5" x="17" y="15.5" rx=".6" transform="rotate(-90 17 15.5)"/><rect width="7" height="5" x="2" y="10" rx=".6" transform="rotate(-90 2 10)"/><path d="M7 17.5h3.5a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2H7m5.5 5.5H17"/></g>`),
		g.Group(children),
	)
}