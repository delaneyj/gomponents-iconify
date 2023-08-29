package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meinblau(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.83 21.77a5.14 5.14 0 1 1 0 10.27h-8.47V11.5h8.47a5.14 5.14 0 1 1 0 10.27Zm0 0h-8.47M16.03 36.5h15.94m-13.61-25h-2.33m2.33 20.54h-2.33"/>`),
		g.Group(children),
	)
}