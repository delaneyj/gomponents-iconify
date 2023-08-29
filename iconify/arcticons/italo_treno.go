package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItaloTreno(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="9.241" cy="29.927" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.789 26.963a3.223 3.223 0 0 1-3.213 3.214h0a3.223 3.223 0 0 1-3.213-3.213v-2.089a3.223 3.223 0 0 1 3.213-3.213h0a3.223 3.223 0 0 1 3.213 3.213m0 5.302v-8.515m3.214-4.339V28.57a1.518 1.518 0 0 0 1.607 1.606h.482M16.15 18.93v11.247m-1.767-8.515h3.374"/><circle cx="11.588" cy="19.18" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.588 21.662v8.515m24.708 0h0a3.223 3.223 0 0 1-3.213-3.214v-2.088a3.223 3.223 0 0 1 3.213-3.213h0a3.223 3.223 0 0 1 3.213 3.213v2.088a3.223 3.223 0 0 1-3.213 3.213Z"/>`),
		g.Group(children),
	)
}