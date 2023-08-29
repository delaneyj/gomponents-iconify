package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33a2 2 0 0 1 2 2v8.33a2 2 0 0 1-2 2H30.17v9.93a2.41 2.41 0 0 1-2.41 2.41h-9.93V40.5a2 2 0 0 1-2 2H7.5a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}