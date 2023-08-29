package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frigate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M11 12v12l13-5l13 5V12H11Zm8-5v5h10V7a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3Z"/><path d="M12 44L6 26l18-7l18 7l-6 18"/><path d="M4 42s4.663 2 8 2c5 0 7-2 12-2s7 2 12 2c3.337 0 8-2 8-2M24 19v23"/></g>`),
		g.Group(children),
	)
}