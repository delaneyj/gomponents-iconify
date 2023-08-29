package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 4h32M8 44h32M12 4v12l9 10m15 18V29.5L27 21M12 44V30l6.5-6.5"/><path d="M36 4v12l-6.5 7.5M18 33h1m10.146-.353l.708.707M24 38h1"/></g>`),
		g.Group(children),
	)
}