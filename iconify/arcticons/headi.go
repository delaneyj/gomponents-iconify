package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.49 18.35a15.34 15.34 0 0 0-30.61 1.49l-3.74 3.74a2.36 2.36 0 0 0 1.66 4h2.08v4.47a5.27 5.27 0 0 0 5.27 5.27h0v6.18h17.59V33.2a15.36 15.36 0 0 0 7.75-14.85Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.53 12.32l-5.17 8h5.17l-3.88 7.57"/>`),
		g.Group(children),
	)
}