package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maingau(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.049 16.962v13.51H38.5m-29 .162h6.426M9.5 17.908h6.426M9.5 24.271h4.177M9.5 17.908v12.726m9.02-1.119a4.541 4.541 0 0 0 3.674 1.523h2.204a3.548 3.548 0 0 0 3.674-3.383h0a3.548 3.548 0 0 0-3.674-3.384H22.01a3.548 3.548 0 0 1-3.674-3.383h0a3.548 3.548 0 0 1 3.674-3.384h2.204a4.085 4.085 0 0 1 3.674 1.523"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}