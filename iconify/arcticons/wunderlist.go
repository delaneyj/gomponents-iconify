package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wunderlist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 7.5v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-6.65V37l-4.92-2.86L24 31.26l-4.93 2.86L14.15 37V5.5H7.5a2 2 0 0 0-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 14.94l2.18 4.41l4.87.71l-3.52 3.44l.83 4.85L24 26.06l-4.36 2.29l.83-4.85L17 20.06l4.87-.71Z"/>`),
		g.Group(children),
	)
}