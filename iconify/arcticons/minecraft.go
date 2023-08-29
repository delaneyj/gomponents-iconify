package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minecraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-4.13 14.12h-8.18V24h4.09v12.27h-4.09v-4.09H20v4.09h-4.08V24H20v-4.38h-8.17v-8.18H20v8.18h8.18v-8.18h8.18Z"/>`),
		g.Group(children),
	)
}