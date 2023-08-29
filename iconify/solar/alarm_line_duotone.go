package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="13" r="9" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v4l2.5 2.5m-11-11l4-2.5m13 2.5l-4-2.5"/></g>`),
		g.Group(children),
	)
}