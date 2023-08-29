package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#5546FF"/><path fill="#FFF" d="m19.319 19.033l3.61 5.467h-2.697l-4.24-6.423l-4.238 6.423H9.07l3.611-5.453H7.5v-2.07h17v2.056zm5.181-6.138v2.085h-17v-2.084h5.081L9.013 7.5h2.698l4.282 6.509L20.289 7.5h2.698l-3.568 5.395z"/></g>`),
		g.Group(children),
	)
}