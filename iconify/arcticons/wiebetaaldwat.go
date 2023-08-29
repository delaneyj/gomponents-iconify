package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wiebetaaldwat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.64 25.14V18.6m0 0h2.75A1.64 1.64 0 0 1 37 20.24h0a1.63 1.63 0 0 1-1.64 1.64h-2.72"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.64 21.88h2.75A1.64 1.64 0 0 1 37 23.52h0a1.63 1.63 0 0 1-1.64 1.64h-2.72m0-16.35l1.64 6.55m1.63-6.55l-1.63 6.55m1.63-6.55l1.63 6.55m1.64-6.55l-1.64 6.55m-4.9 13.05l1.64 6.55m1.63-6.55l-1.63 6.55m1.63-6.55l1.63 6.55m1.64-6.55l-1.64 6.55"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM8.82 38.43h30.36"/>`),
		g.Group(children),
	)
}