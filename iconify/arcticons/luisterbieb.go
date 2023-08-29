package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luisterbieb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.35 5.6c6.12 0 10.15 3.21 10.15 10.81s-4.08 26-17.8 26S4.5 29.2 4.5 22.87C4.5 13.13 23.62 5.6 33.35 5.6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.27 23.58h.84a1 1 0 0 1 1 1v6.92a1 1 0 0 1-1 1h-.84a4.43 4.43 0 0 1-4.44-4.5v0a4.43 4.43 0 0 1 4.44-4.42Zm15.46-.3a7.73 7.73 0 0 0-15.46 0m0 .3v8.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.74 32.16h-.84a1 1 0 0 1-1-1v-6.92a1 1 0 0 1 1-1h.84a4.43 4.43 0 0 1 4.43 4.43h0a4.43 4.43 0 0 1-4.43 4.49Zm-.01-8.88v8.87"/>`),
		g.Group(children),
	)
}