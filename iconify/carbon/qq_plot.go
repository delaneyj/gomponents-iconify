package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QqPlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="20" cy="4" r="2" fill="currentColor"/><circle cx="8" cy="16" r="2" fill="currentColor"/><circle cx="28" cy="12" r="2" fill="currentColor"/><circle cx="11" cy="7" r="2" fill="currentColor"/><circle cx="16" cy="24" r="2" fill="currentColor"/><path fill="currentColor" d="M30 3.413L28.586 2L4 26.585V2H2v26a2 2 0 0 0 2 2h26v-2H5.413Z"/>`),
		g.Group(children),
	)
}