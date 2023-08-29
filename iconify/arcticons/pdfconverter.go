package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pdfconverter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.2" height="39" x="8.4" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28 27.32l7.4 10.73a.56.56 0 0 1-.45.87h-21.9a.56.56 0 0 1-.39-1l6.14-6.13a.54.54 0 0 1 .78 0l1.09 1.08a.54.54 0 0 0 .78 0l5.68-5.68a.55.55 0 0 1 .87.13ZM19.55 9.08H12.5v11.15l3.53-3.03l3.52 3.03V9.08z"/>`),
		g.Group(children),
	)
}