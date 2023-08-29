package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binauralbeats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.68 25.54a4.42 4.42 0 0 1 4.41-4.41h0a4.41 4.41 0 0 1 4.41 4.41v2.87a4.41 4.41 0 0 1-4.41 4.41h0a4.42 4.42 0 0 1-4.41-4.41m0 4.41V15.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M18.77 24a4.41 4.41 0 0 1 0 8.82H11.5V15.18h7.27a4.41 4.41 0 0 1 0 8.82Zm0 0H11.5"/>`),
		g.Group(children),
	)
}