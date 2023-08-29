package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenBanknote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M2 10v44h60V10H2m37.273 24.038a7.165 7.165 0 0 0 4.545 1.629c4.018 0 7.271-3.284 7.271-7.334S47.836 21 43.818 21a7.17 7.17 0 0 0-4.545 1.628v-5.295h15.453v22H39.273v-5.295m-14.546 5.295H9.273v-22h15.454v22m0-24.75H6.545v27.5h18.182v2.75H3.818v-33h20.909v2.75m12.728 30.253h-10.91V12.75h10.91v32.086m22.727-.003H39.273v-2.75h18.182v-27.5H39.273v-2.75h20.908v33z"/><path fill="currentColor" d="M13.985 26.056v2.034h2.603v2.034h-2.603v2.032h2.603v3.051h1.733v-3.051h2.601v-2.032h-2.601V28.09h2.601v-2.034H18.82l4.089-3.945l-1.102-1.569l-4.352 4.196l-4.352-4.196L12 22.111l4.089 3.945z"/>`),
		g.Group(children),
	)
}