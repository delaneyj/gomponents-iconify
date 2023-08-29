package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cxxdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.642 21.288l-4 5.3m-9.819.239h0a5.713 5.713 0 0 1-5.816 5.653h0a5.713 5.713 0 0 1-5.816-5.653v-5.654a5.713 5.713 0 0 1 5.816-5.653h0a5.53 5.53 0 0 1 5.6 5.653h0m6.98.039l-4 5.3m4 0l-4-5.3m11.222 5.3l-4-5.3"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.764 4.373v8.555a2 2 0 0 0 2 2h8.724A21.5 21.5 0 0 0 32.764 4.373Z"/>`),
		g.Group(children),
	)
}