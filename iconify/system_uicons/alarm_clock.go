package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5 4.565h-2a6 6 0 0 0-6 6V12.5a6 6 0 0 0 6 6h2a6 6 0 0 0 6-6v-1.935a6 6 0 0 0-6-6zm3.032-1.068c.884-.639 2.089-.71 2.968.003c.906.734 1.258 1.96.822 2.969M6.532 3.544C5.642 2.862 4.4 2.77 3.5 3.5c-.906.734-1.258 1.96-.822 2.97"/><path d="M10.5 7.5v4H14M5 17l-2 2m13-2l2 2"/></g>`),
		g.Group(children),
	)
}