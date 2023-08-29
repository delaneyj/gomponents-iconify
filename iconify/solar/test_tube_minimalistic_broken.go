package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m6.8 11.783l1.275.143a2.205 2.205 0 0 1 1.944 1.952a2.209 2.209 0 0 0 1.32 1.787l1.661.69m0 0l-3.512 3.527a3.789 3.789 0 0 1-5.375 0a3.83 3.83 0 0 1 0-5.4l10.75-10.797l5.376 5.399l-1.81 1.818M13 16.355l3-3.014m5-3.492L14.181 3"/>`),
		g.Group(children),
	)
}