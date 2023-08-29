package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M10.25 50.857h10.625v4.231H10.25z"/><path fill="#D0CFCE" d="M51 17.084h10.625v4.231H51z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M21 53.276h35.564s6.436 0 6.436-9.193s-6.115-9.193-6.115-9.193H17.577s-5.85 0-5.85-8.303s5.85-7.354 5.85-7.354h33.219"/><path d="M10.25 50.857h10.625v4.231H10.25zM51 17.084h10.625v4.231H51z"/></g>`),
		g.Group(children),
	)
}