package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pleco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.073 20.34C18.685 16 23.085 11.792 24.78 7.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.333 13.871l10.743-1.617l-9.01 9.242l-7.163.52l3.35 11.494l13.054-1.733l2.715-11.898l-9.3 1.213l-.692 11.552m-4.16-5.545l9.011-1.387m-17.617 9.415L7.451 42.5m9.069-5.928l3.582 3.985m4.909-4.274l3.582 4.332m5.372-5.199l6.584 6.758M24.723 21.091l-2.657.405"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}