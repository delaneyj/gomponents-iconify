package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openmoneybox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.7 23.25v-7.12a2 2 0 0 1 2-2h32.8a2 2 0 0 1 2 2v23a2 2 0 0 1-2 2H8.7a2 2 0 0 1-2-2V32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 23.25H17a2 2 0 0 1 2 2V30a2 2 0 0 1-2 2H5.5a1 1 0 0 1-1-1v-6.75a1 1 0 0 1 1-1Z"/><circle cx="14.49" cy="27.63" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.12 14.13L13.19 6.87l-2.11 7.26"/>`),
		g.Group(children),
	)
}