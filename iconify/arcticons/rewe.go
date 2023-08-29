package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.459 28.166h4.166m-4.166-8.332h4.166M17.459 24h2.708m-2.708-4.166v8.332m16.875 0H38.5m-4.166-8.332H38.5M34.334 24h2.708m-2.708-4.166v8.332m-24.834 0v-8.332h2.708a2.812 2.812 0 1 1 .002 5.624H9.5m2.906-.007l2.614 2.715m17.055-8.332l-2.083 8.332l-2.083-8.332l-2.082 8.332l-2.083-8.332"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.066 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}