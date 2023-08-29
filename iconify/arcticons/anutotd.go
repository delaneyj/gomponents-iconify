package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anutotd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2.006 2.006 0 0 0-2 2v33a2.006 2.006 0 0 0 2 2h33a2.006 2.006 0 0 0 2-2v-33a2.006 2.006 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.7 16.5v15.1h-7.5v-15A11.032 11.032 0 1 0 35 27a10.817 10.817 0 0 0-7.3-10.5Zm-7.5-3.1h7.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.2 9.1h7.5v22.4h-7.5z"/>`),
		g.Group(children),
	)
}