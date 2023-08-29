package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covariate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="8" cy="16" r="2" fill="currentColor"/><circle cx="14" cy="8" r="2" fill="currentColor"/><circle cx="28" cy="12" r="2" fill="currentColor"/><circle cx="21" cy="18" r="2" fill="currentColor"/><path fill="currentColor" d="M30 3.414L28.586 2L4 26.586V2H2v26a2 2 0 0 0 2 2h26v-2H5.414ZM4 28Z"/>`),
		g.Group(children),
	)
}