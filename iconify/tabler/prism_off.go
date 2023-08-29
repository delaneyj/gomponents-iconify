package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrismOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12v10m5.957-4.048l-4.937 3.703a1.7 1.7 0 0 1-2.04 0L5 17.17a2.5 2.5 0 0 1-1-2V4m3-1h12a1 1 0 0 1 1 1v11.17c0 .25-.037.495-.109.729"/><path d="M12.688 8.7a1.7 1.7 0 0 0 .357-.214L19.7 3.3M3 3l18 18"/></g>`),
		g.Group(children),
	)
}