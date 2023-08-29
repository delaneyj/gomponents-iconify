package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 32.5v5a5 5 0 0 1-5 5h0a5 5 0 0 1-5-5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 37.5a5 5 0 0 1 5-5h25c2.2 0 4-1.8 4-4v-19c0-2.2-1.8-4-4-4H11c-2.2 0-4 1.8-4 4v28"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.188 18.467l-9.145-5.265a1.233 1.233 0 0 0-1.848 1.068V24.8a1.233 1.233 0 0 0 1.848 1.07l9.145-5.266a1.233 1.233 0 0 0 0-2.137Z"/>`),
		g.Group(children),
	)
}