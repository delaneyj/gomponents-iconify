package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudtvRemote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.343 39.062a21.5 21.5 0 1 1-.14-30.265"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.127 24l-16.672-9.625v19.25L34.127 24z"/>`),
		g.Group(children),
	)
}