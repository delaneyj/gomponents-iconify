package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Networkmapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.2 8.8A21.5 21.5 0 1 1 24 2.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.5 24a3.5 3.5 0 1 1-3.5-3.5a3.5 3.5 0 0 1 3.5 3.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.7 15.3a12.3 12.3 0 1 1-8.7-3.6m0-9.2v18"/>`),
		g.Group(children),
	)
}