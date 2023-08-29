package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartIncreasingWithYen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-11 8h6.078l4.894 4.887L36.896 10H43l-8.881 8.813v.006h4.6v4.297h-4.6v4.295h4.6v4.295h-4.6V36h-4.302v-4.295H25.3V27.41h4.517v-4.295H25.3v-4.297h4.517v-.01L21 10.004V10zm30.322 28.986l-3.977-2.707l-6.555 11.475l-4.855-2.689l-4.358 8.199l-5.07-3.637L23.271 54A26.504 26.504 0 0 1 19 52.23l6.482-8.762l4.478 3.209l4.081-7.672l4.977 2.756l6.777-11.863l4.256 2.895L55.92 24A22.945 22.945 0 0 1 57 30.48l-5.678 8.506z"/>`),
		g.Group(children),
	)
}