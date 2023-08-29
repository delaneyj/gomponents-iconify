package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouserEiuMagazine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.739 31.113V8.5L24 31.139L35.261 8.534v22.605M17.391 39.5h3.057m-3.057-6.114h3.057m-3.057 3.057h1.987m-1.987-3.057V39.5m5.202 0h1.834m-1.834-6.114h1.834m-.917 0V39.5m3.125-6.114v4.05a1.988 1.988 0 1 0 3.974 0v-4.05"/>`),
		g.Group(children),
	)
}