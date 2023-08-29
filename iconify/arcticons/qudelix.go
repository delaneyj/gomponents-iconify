package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qudelix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.845 21.462c0 3.988-3.53 7.25-7.845 7.25s-7.845-3.262-7.845-7.25V16.75c0-3.987 3.53-7.25 7.845-7.25s7.845 3.263 7.845 7.25m0-7.25v29" class="b"/>`),
		g.Group(children),
	)
}