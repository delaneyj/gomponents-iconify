package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ytdlnis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.635 18.09a4.88 4.88 0 0 0-3.445-3.446c-2.257-.87-24.114-1.298-30.455.025a4.88 4.88 0 0 0-3.445 3.446c-1.019 4.468-1.096 14.129.025 18.697a4.881 4.881 0 0 0 3.445 3.445c4.468 1.028 25.712 1.173 30.455 0a4.88 4.88 0 0 0 3.445-3.445c1.086-4.868 1.163-13.93-.025-18.723Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 32.543l5.841-10.185H18.16L24 32.543Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 22.358v-8.963a6.479 6.479 0 1 1 12.957 0v.891"/>`),
		g.Group(children),
	)
}