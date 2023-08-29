package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Themes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v1.89h37V7.45a2 2 0 0 0-2-1.95Zm3.89 7.3a2.92 2.92 0 0 0-2.92 2.92v6.33A2.91 2.91 0 0 1 5.5 25v15.55a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V23.24a2.88 2.88 0 0 1-2 .76a2.92 2.92 0 0 1-2.92-2.93v-5.35a2.92 2.92 0 0 0-5.84 0v1.46a2.92 2.92 0 0 1-5.84 0v-1.46a2.92 2.92 0 1 0-5.84 0V25a2.93 2.93 0 0 1-5.85 0v-9.28a2.91 2.91 0 0 0-2.92-2.92Z"/>`),
		g.Group(children),
	)
}