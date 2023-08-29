package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bigbasket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-28.777 6.139V36.39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.072 28.11a7.85 7.85 0 0 1-7.62 8.034a7.84 7.84 0 0 1-7.713-7.936a7.856 7.856 0 0 1 7.565-8.09a7.833 7.833 0 0 1 7.767 7.878M20.838 11.61v7.162"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.613 20.953a7.582 7.582 0 0 1 9.06 11.629a8.163 8.163 0 0 1-9.119 2.946"/>`),
		g.Group(children),
	)
}