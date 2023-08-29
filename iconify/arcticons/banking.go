package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.646 24.1a1.719 1.719 0 0 1 0 3.438H17.81v-6.875h2.836a1.719 1.719 0 0 1 0 3.437Zm0 0H17.81m-5.631 1.06a2.32 2.32 0 0 1-4.64 0v-2.32a2.31 2.31 0 0 1 2.32-2.32a2.24 2.24 0 0 1 2.234 2.32m16.92 2.32a2.32 2.32 0 0 1-4.64 0v-2.32a2.31 2.31 0 0 1 2.32-2.32a2.24 2.24 0 0 1 2.234 2.32m-14.909 4.641h2.062m-2.062-6.876h2.062m-1.031 0v6.876m21.014-5.117L34.689 24l1.374 1.636m0-3.272L37.437 24l-1.374 1.636"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.458 20.305h-2.791L31.581 24l3.086 3.696h2.79L40.547 24Z"/>`),
		g.Group(children),
	)
}