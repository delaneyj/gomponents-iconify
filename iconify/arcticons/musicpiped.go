package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicpiped(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.51 21.51 0 0 0 2.5 24h0A21.51 21.51 0 0 0 24 45.5h0a21.5 21.5 0 0 0 0-43Z"/><circle cx="21.49" cy="30.86" r="5.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.13 30.86V11.5h5.01"/>`),
		g.Group(children),
	)
}