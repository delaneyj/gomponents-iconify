package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Network(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="7" height="5" x="3" y="2" rx=".6"/><rect width="7" height="5" x="8.5" y="17" rx=".6"/><rect width="7" height="5" x="14" y="2" rx=".6"/><path d="M6.5 7v3.5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V7M12 12.5V17"/></g>`),
		g.Group(children),
	)
}