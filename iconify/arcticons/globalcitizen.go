package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globalcitizen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.503 21.503 0 0 1 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M24 11.5A12.5 12.5 0 1 1 11.5 24A12.502 12.502 0 0 1 24 11.5Z"/>`),
		g.Group(children),
	)
}