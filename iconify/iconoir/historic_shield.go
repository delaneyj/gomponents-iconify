package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoricShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 15.528V2.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6v12.928a4 4 0 0 1-2.211 3.578l-5.52 2.76a.6.6 0 0 1-.537 0l-5.52-2.76A4 4 0 0 1 4 15.528Z"/>`),
		g.Group(children),
	)
}