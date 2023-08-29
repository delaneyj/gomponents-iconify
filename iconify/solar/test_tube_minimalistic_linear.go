package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m6.8 11.783l1.275.143a2.205 2.205 0 0 1 1.944 1.952a2.209 2.209 0 0 0 1.32 1.787l1.661.69m0 0l7.239-7.271l-5.376-5.399l-10.75 10.798a3.83 3.83 0 0 0 0 5.399a3.789 3.789 0 0 0 5.375 0L13 16.355Zm8-6.506L14.181 3"/>`),
		g.Group(children),
	)
}