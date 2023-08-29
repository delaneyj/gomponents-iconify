package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.444 15.593a16.311 16.311 0 1 1-22.987 0L23.95 4.475Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.295 22.033a7.584 7.584 0 1 1-10.689 0l5.345-5.17Z"/>`),
		g.Group(children),
	)
}