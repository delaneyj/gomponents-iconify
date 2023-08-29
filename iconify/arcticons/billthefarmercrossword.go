package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Billthefarmercrossword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.833 5.5v37m12.334-37v37M5.5 17.833h37m-37 12.334h37m-32.833-14.5h4m-4-6.911l2-1.089m0 0v8m22.016-5.35a2.65 2.65 0 1 1 5.3 0a2.473 2.473 0 0 1-.776 1.874c-1.072.94-4.524 3.476-4.524 3.476h5.3M12.317 36.334a2 2 0 0 0 2-2h0a2 2 0 0 0-2-2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.317 40.333a2 2 0 0 0 2-2h0a2 2 0 0 0-2-2m-3.299 3.325a3.407 3.407 0 0 0 2.487.675h.812m-3.3-7.331a3.407 3.407 0 0 1 2.49-.669l.811.002m-2.038 3.999h2.037"/>`),
		g.Group(children),
	)
}