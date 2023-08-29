package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrideHealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.67 12.962c2.046-1.98 2.388-3.14 2.388-5.084S36.966 4.5 33.758 4.5S17.892 8.185 17.892 15.487s15.935 11.26 15.935 18.289S24.409 43.5 19.018 43.5s-9.076-2.696-9.076-6.756s7.711-7.063 7.711-7.063"/>`),
		g.Group(children),
	)
}