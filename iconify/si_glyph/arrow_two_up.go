package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTwoUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.973 5.175l5.002 3.744v-3.94L8.973 1.036L4.004 5.078v3.918l.012.011l4.957-3.832zm0 7.783l5.002 3.951v-3.938L8.973 9.005l-4.969 4.064v3.918l.012.01l4.957-4.039z"/>`),
		g.Group(children),
	)
}