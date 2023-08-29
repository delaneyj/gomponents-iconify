package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#000" fill-rule="nonzero"/><path stroke="#FFF" stroke-linecap="round" stroke-linejoin="round" stroke-width=".64" d="M10.886 11.61L16 27.667l-7.588-4.754l2.474-11.303L16 4.624v4.9L8.412 22.913h15.183L16.007 9.524v-4.9l5.113 6.986l2.475 11.303l-7.588 4.754L21.12 11.61"/></g>`),
		g.Group(children),
	)
}