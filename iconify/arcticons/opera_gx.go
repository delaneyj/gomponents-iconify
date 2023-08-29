package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OperaGx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.519 5.16A18.839 18.839 0 0 0 12.68 24h0a18.647 18.647 0 0 0 22.185 18.501A21.45 21.45 0 0 0 34.9 5.473a18.84 18.84 0 0 0-3.38-.312Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.624 9.362a18.841 18.841 0 0 1-.012 29.286"/>`),
		g.Group(children),
	)
}