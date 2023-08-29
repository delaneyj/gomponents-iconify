package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Her(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM11.65 19.319v10.362M18.169 19v11m-7.288-5.5h7.288m3.712 5.5h5.5m-5.5-11h5.5m-5.5 3.781h3.575m-3.575 3.438h3.575M21.881 19v11"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.131 30V19h3.575a3.712 3.712 0 1 1 0 7.425h-3.575m3.839-.012L37.419 30"/>`),
		g.Group(children),
	)
}