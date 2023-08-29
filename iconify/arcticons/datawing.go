package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Datawing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5L15.5 30l8.87-2.81l8.13 2.79Zm-4.9 24.32V43.5m3.23-15.71V43.5m3.86-15.73V43.5M29.63 29v14.5"/>`),
		g.Group(children),
	)
}