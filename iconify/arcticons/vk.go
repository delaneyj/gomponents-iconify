package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.55 35.19v-6.64c4.46.68 5.87 4.19 8.71 6.64h7.24a29.36 29.36 0 0 0-7.9-10.47c2.6-3.58 5.36-6.95 6.71-12.06h-6.58c-2.58 3.91-3.94 8.49-8.18 11.51V12.66H18l2.28 2.82v10.05c-3.7-.43-6.2-7.2-8.91-12.87H4.5c2.5 7.66 7.76 24.47 23.05 22.53Z"/>`),
		g.Group(children),
	)
}