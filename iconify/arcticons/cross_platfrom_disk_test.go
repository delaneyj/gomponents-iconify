package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossPlatfromDiskTest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.24 39V25.751h2.981a5.796 5.796 0 0 1 5.796 5.796v1.657A5.796 5.796 0 0 1 16.221 39Zm12.742-16.751V9h4.338a4.45 4.45 0 0 1 0 8.9h-4.338m0 7.851h8.778M30.371 39V25.751m-8.353-7.945v.054a4.389 4.389 0 0 1-4.389 4.389h0a4.389 4.389 0 0 1-4.389-4.389v-4.471A4.389 4.389 0 0 1 17.629 9h0a4.389 4.389 0 0 1 4.389 4.389v.054"/>`),
		g.Group(children),
	)
}