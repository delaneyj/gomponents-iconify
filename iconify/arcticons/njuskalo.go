package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Njuskalo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.754 19.102v11.3c0 2.08-1.67 3.767-3.732 3.767h0a3.702 3.702 0 0 1-2.638-1.103"/><circle cx="30.754" cy="14.581" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.03 29.084v-6.216a3.767 3.767 0 0 0-3.767-3.767h0a3.767 3.767 0 0 0-3.767 3.767m0 6.216v-9.982"/>`),
		g.Group(children),
	)
}