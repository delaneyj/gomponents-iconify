package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerkehrNrw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.558 37.492v-6.406h2.09c1.206 0 2.171.96 2.171 2.162s-.965 2.162-2.17 2.162h-2.091m2.243-.006l2.018 2.088M11.77 29.539h3.22m-3.22-6.358h3.22m-3.22 3.179h2.093m-2.093-3.179v6.358m4.63 0V23.18h2.092c1.208 0 2.174.953 2.174 2.146s-.966 2.145-2.174 2.145H16.4m2.246-.005l2.02 2.073m1.409-6.358v6.358m3.461 0l-2.656-3.18l2.656-3.178M22.88 26.36h-.805m9.499-3.179v6.358m4.267-6.358v6.358m-4.267-3.179h4.267m-25.48-3.18l-2.093 6.36l-2.173-6.358m20.85 6.357h3.22m-3.22-6.358h3.22m-3.22 3.179h2.093m-2.093-3.179v6.358m10.305 0V23.18h2.092c1.208 0 2.174.953 2.174 2.146s-.966 2.145-2.174 2.145H37.25m2.246-.005l2.02 2.073M23.53 37.492v-6.406l4.262 6.406v-6.406m13.726 0l-1.608 6.406l-1.608-6.406l-1.608 6.406l-1.608-6.406"/><circle cx="21.718" cy="37.251" r=".753" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 10.5v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2Z"/>`),
		g.Group(children),
	)
}