package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 13.01l.01-.011M7 17h10M6 10V3.6a.6.6 0 0 1 .6-.6h10.8a.6.6 0 0 1 .6.6V10m3 10.4V14a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v6.4a.6.6 0 0 0 .6.6h16.8a.6.6 0 0 0 .6-.6Z"/>`),
		g.Group(children),
	)
}