package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoodLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM5.5 24h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 34.31h-.1a4.502 4.502 0 0 1-8.9-.944a4.5 4.5 0 0 1 8.9-.945h.1M24 42.5v-8.189m0-1.891V24m0-10.31h.1a4.502 4.502 0 0 1 8.9.944a4.5 4.5 0 0 1-8.9.945H24M24 5.5v8.189m0 1.89V24"/>`),
		g.Group(children),
	)
}