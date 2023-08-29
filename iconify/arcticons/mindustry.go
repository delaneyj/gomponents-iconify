package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mindustry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.65 4.5L4.51 16.65v14.7L7 33.83L15.81 25H32.2l8.8 8.83l2.47-2.48v-14.7L31.35 4.5H28v11.41l-2.31 2.31h-3.35L20 15.91V4.5Zm2.19 26.65l-7.27 7.27l5.09 5.08h14.68l5.09-5.09l-7.26-7.26H18.84Z"/>`),
		g.Group(children),
	)
}