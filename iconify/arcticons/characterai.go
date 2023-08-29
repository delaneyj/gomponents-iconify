package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Characterai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.282 25.897c-.45.9-1.5 1.5-2.55 1.5h0a3.01 3.01 0 0 1-3-3v-1.95a3.01 3.01 0 0 1 3-3h0c1.05 0 2.1.6 2.55 1.5"/><circle cx="33.518" cy="16.853" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.518 19.297v8.1m-3.558 0v-8.1m0 5.1a3.01 3.01 0 0 1-3 3h0a3.01 3.01 0 0 1-3-3v-1.95a3.01 3.01 0 0 1 3-3h0a3.01 3.01 0 0 1 3 3"/><circle cx="21.879" cy="27.147" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}