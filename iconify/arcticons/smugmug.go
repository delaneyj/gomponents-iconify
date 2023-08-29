package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smugmug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.72 22.643c1.038-1.687 7.667-.391 14.774-.714c6.968-.317 13.52-1.947 14.817-.26c1.626 2.115-6.62 20.831-19.858 20.831S7.886 25.622 9.72 22.643Zm6.813-16.705c-1.817 0-4.023 1.428-4.023 3.05s.973 2.726 3.115 2.726s4.607-.844 4.607-2.888s-1.622-2.888-3.699-2.888ZM35.086 5.5c-1.817 0-4.023 1.428-4.023 3.05s.973 2.725 3.115 2.725s4.607-.843 4.607-2.887S37.163 5.5 35.086 5.5Z"/>`),
		g.Group(children),
	)
}