package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panasonicapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.643 24.171h4.65v9.767h7.797a6.325 6.325 0 0 1-2.957-7.115A6.262 6.262 0 0 1 24 22.13a6.262 6.262 0 0 1 5.867 4.693a6.325 6.325 0 0 1-2.958 7.115h7.797v-9.767h4.65L24 14.061Z"/><circle cx="24" cy="28.404" r="2.887" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}