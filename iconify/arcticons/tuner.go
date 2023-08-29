package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm10.66 0v37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.16 9.89c3.7 0 2.24 28.92 9.92 28.92H42.5M18.16 9.89c-3.7 0-2.24 28.92-9.93 28.92H5.5"/>`),
		g.Group(children),
	)
}