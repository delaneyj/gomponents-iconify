package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDArcCenterPt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-dasharray="3 3" d="M22 16c0-5.523-4.477-10-10-10c-4.1 0-7.625 2.468-9.168 6"/><path fill="currentColor" d="M2 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M2 16h10"/><path fill="currentColor" d="M12 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`),
		g.Group(children),
	)
}