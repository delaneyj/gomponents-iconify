package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vipps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="28.304" cy="16.948" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.083 23.803c5.977 8.485 14.174 11.149 21.577 1.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}