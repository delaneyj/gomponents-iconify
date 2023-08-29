package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Entry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21 5.387a1 1 0 0 1 1.316-.948l13 4.333a1 1 0 0 1 .684.949v28.558a1 1 0 0 1-.684.949l-13 4.333A1 1 0 0 1 21 42.613V39h-9V8h9V5.387ZM27 23c0 1.105-.448 2-1 2s-1-.895-1-2s.448-2 1-2s1 .895 1 2Zm-6-13h-7v27h7V10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}