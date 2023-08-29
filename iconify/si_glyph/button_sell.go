package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonSell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.343 2.033L9.01 3.083l-5.375-1.05C2.18 2.033 1 3.359 1 4.996v5.018c0 1.637 1.18 2.964 2.635 2.964l5.375-1.02l5.333 1.02c1.456 0 2.636-1.327 2.636-2.964V4.996c0-1.637-1.18-2.963-2.636-2.963zM6.042 6.004H4.021v.97h.995v1.009h1.026V9.08h-1v.978H2.98V8.939h1.989v-.907H3.98V7.023h-1V5.958h1V4.949h2.062v1.055zm3 .026H8.021v.962h.995v1.051h-.974v.873h1v1.141H6.98V4.979h2.062V6.03zm3 4.027H9.984V4.958h1.031v3.979h1.026v1.12h.001zm3 0h-2.058V4.948h1.031V9h1.026v1.057h.001z"/>`),
		g.Group(children),
	)
}