package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IconsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.01 4.041A3.5 3.5 0 0 0 6.5 10c.975 0 1.865-.357 2.5-1m.958-3.044a3.503 3.503 0 0 0-2.905-2.912M2.5 21h8l-4-7zM14 3l7 7m-7 0l7-7m-3 11h3v3m0 4h-7v-7M3 3l18 18"/>`),
		g.Group(children),
	)
}