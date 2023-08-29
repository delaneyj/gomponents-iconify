package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderEpicene(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.536 15.536a5 5 0 1 0-7.072-7.072a5 5 0 0 0 7.072 7.072zm0-.001L21 10M3 14l5.464-5.535M12 12h.01"/>`),
		g.Group(children),
	)
}