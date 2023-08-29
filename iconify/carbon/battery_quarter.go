package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryQuarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 24H6a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h18a2.002 2.002 0 0 1 2 2v1h1a2.002 2.002 0 0 1 2 2v6a2.003 2.003 0 0 1-2 2h-1v1a2.003 2.003 0 0 1-2 2ZM6 10v12h18v-3h3v-6h-3v-3Z"/><path fill="currentColor" d="M12 12v8H8v-8z"/>`),
		g.Group(children),
	)
}