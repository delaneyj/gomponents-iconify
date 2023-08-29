package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwiPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-5.659 4.272v10.559"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.117 9.772l-2.64 10.559l-2.639-10.559l-2.64 10.559l-2.64-10.559m-9.078 9.402a2.953 2.953 0 0 0 2.589 1.157h1.563a2.637 2.637 0 0 0 2.634-2.64h0a2.637 2.637 0 0 0-2.634-2.64h-1.727a2.637 2.637 0 0 1-2.634-2.64h0a2.637 2.637 0 0 1 2.634-2.64h1.563a2.954 2.954 0 0 1 2.589 1.157M5.5 24h37m-18.917 6.056v3.69a2.236 2.236 0 0 0 2.235 2.235h0a2.236 2.236 0 0 0 2.236-2.236v-3.688m0 3.688v2.236m2.411-.5a2.515 2.515 0 0 0 1.838.5h.502a1.48 1.48 0 0 0 1.478-1.481h0a1.48 1.48 0 0 0-1.478-1.481h-1.003a1.48 1.48 0 0 1-1.478-1.481h0a1.48 1.48 0 0 1 1.478-1.482h.501a2.514 2.514 0 0 1 1.839.5m-13.687-3.518v7.825a1.118 1.118 0 0 0 1.118 1.118h.335m-8.191-2.236a2.236 2.236 0 0 0 2.236 2.236h0a2.236 2.236 0 0 0 2.236-2.236v-1.453a2.236 2.236 0 0 0-2.236-2.236h0a2.236 2.236 0 0 0-2.236 2.236m0-2.236V39"/>`),
		g.Group(children),
	)
}