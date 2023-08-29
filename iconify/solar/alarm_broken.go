package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 9v4l2.5 2.5m-11-11l4-2.5m13 2.5l-4-2.5"/><path d="M7.5 5.204A9 9 0 1 1 4.204 8.5"/></g>`),
		g.Group(children),
	)
}