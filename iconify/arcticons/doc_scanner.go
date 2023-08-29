package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM5.5 31h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 22.873v-9.246h2.08a4.045 4.045 0 0 1 4.045 4.045v1.156a4.045 4.045 0 0 1-4.045 4.045Zm23-3.101v.038a3.063 3.063 0 0 1-3.063 3.063h0a3.063 3.063 0 0 1-3.063-3.063v-3.12a3.063 3.063 0 0 1 3.063-3.063h0A3.063 3.063 0 0 1 35.5 16.69v.038M20.937 19.81a3.063 3.063 0 1 0 6.126 0v-3.12a3.063 3.063 0 1 0-6.126 0Z"/>`),
		g.Group(children),
	)
}