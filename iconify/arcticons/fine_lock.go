package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FineLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM5.5 24h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24v7.907h.002a3.588 3.588 0 0 1 3.922-2.19a3.143 3.143 0 0 1-.007 6.28A3.586 3.586 0 0 1 24 33.8v8.7m0-37v7.907h-.002a3.588 3.588 0 0 0-3.922-2.19a3.307 3.307 0 0 0-3.009 3.143a3.31 3.31 0 0 0 3.016 3.137A3.586 3.586 0 0 0 24 15.3V24"/>`),
		g.Group(children),
	)
}