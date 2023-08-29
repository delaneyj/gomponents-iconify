package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stronglifts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.438 22.095l-2.876 3.81m2.876 0l-2.876-3.81m9.324-18.167a21.639 21.639 0 0 1 13.195 24.078c-2.193 14.707-21.224 22.619-33.196 13.8c-12.666-7.79-12.471-28.4.34-35.948l3.774-1.944"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.388 28.936c.897.748 1.645 1.048 3.59 1.048h.448a3.96 3.96 0 0 0 3.89-3.89h0a3.96 3.96 0 0 0-3.89-3.89h-4.039v-4.188h7.929m7.441 10.92c.898.748 1.645 1.048 3.59 1.048h.449a3.96 3.96 0 0 0 3.89-3.89h0a3.96 3.96 0 0 0-3.89-3.89h-4.039v-4.188h7.928"/>`),
		g.Group(children),
	)
}