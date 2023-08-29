package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autonavi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.011 30.09L5.5 23.01l37-16.394l-8.736 34.768l-15.51-10.268L42.5 6.616Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.011 30.09l.865 9.295l2.378-8.269"/>`),
		g.Group(children),
	)
}