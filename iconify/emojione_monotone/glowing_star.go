package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlowingStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M62 23.012H39.082L32 2l-7.08 21.012H2L20.541 36l-7.082 21.01L32 44.023L50.541 57.01L43.459 36L62 23.012z"/><path fill="currentColor" d="M46.234 20.344L50.16 8.857l-10.439 7.211l1.461 4.276zM27.912 50.035L32 62l4.09-11.965L32 47.211zM50.68 34.307l-3.825 2.642l1.624 4.752h12.904zM24.277 16.068L13.84 8.857l3.926 11.487h5.052zM13.32 34.307L2.617 41.701H15.52l1.625-4.752z"/>`),
		g.Group(children),
	)
}