package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M38.3 23.47L34.92 20a1.14 1.14 0 0 0-1.63-.23a34.74 34.74 0 0 1-3 2.05a1.81 1.81 0 0 0-1 1.44v2.92A16.58 16.58 0 0 1 24 27a17.18 17.18 0 0 1-5.27-.93v-2.91a1.8 1.8 0 0 0-1-1.46a34.91 34.91 0 0 1-3-2.12a1.15 1.15 0 0 0-1.64.19l-3.39 3.4a.85.85 0 0 0 0 1.1c2.82 3.35 7.84 5 8.82 5.27l.15.05a20.23 20.23 0 0 0 10.45.1s5.9-1.48 9.12-5.13a.86.86 0 0 0 .06-1.09Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.19 34.26A17.5 17.5 0 0 0 16.62 8.12L7.45 5.5l2.35 8.25a17.5 17.5 0 0 0 21.58 26.13l9.16 2.62Z"/>`),
		g.Group(children),
	)
}