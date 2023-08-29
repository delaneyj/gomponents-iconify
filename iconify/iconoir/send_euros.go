package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendEuros(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 12h8m0 0l-3.84-4M21 12l-3.84 4M11 7.503A4.746 4.746 0 0 0 8.87 7C6.18 7 4 9.239 4 12s2.18 5 4.87 5a4.73 4.73 0 0 0 2.13-.503M3 11h6m-6 2h6"/>`),
		g.Group(children),
	)
}