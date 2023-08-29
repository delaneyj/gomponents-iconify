package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16M9 5.336c4.143-.94 8.643.094 12 3.101m-8 .615a9.456 9.456 0 0 1 5.333 2.367m-10 2.981a5.491 5.491 0 0 1 4.345-1.358M3 8.437a13.445 13.445 0 0 1 3.206-2.134m-.54 5.116A9.446 9.446 0 0 1 9 9.484m3 9.016l1-1.118a1.5 1.5 0 0 0-2 0l1 1.118Z"/>`),
		g.Group(children),
	)
}