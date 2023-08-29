package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oeffidirections(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.81 36.9V22.56a6 6 0 0 1 6-6h4.92"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.19 16.56l-4.73 2.74l-4.73 2.73V11.1l9.46 5.46z"/>`),
		g.Group(children),
	)
}