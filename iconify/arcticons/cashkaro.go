package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cashkaro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.984 27.352v.083a6.742 6.742 0 0 1-6.742 6.742h0A6.742 6.742 0 0 1 9.5 27.435v-6.87a6.742 6.742 0 0 1 6.742-6.741h0a6.742 6.742 0 0 1 6.742 6.742v.083m4.576-6.826v20.354m0-7.09L38.5 13.892m0 20.285L30.12 24"/>`),
		g.Group(children),
	)
}