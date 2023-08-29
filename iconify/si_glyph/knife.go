package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.344 5.262L2.746 2.664c-.454-.453-.943.032-1.525-.551C.639 1.531 1.294.745 1.7.34c.406-.405 1.025-.444 1.383-.087l3.275 3.276c.357.358-.657 2.092-1.014 1.733zm11.594 10.363c-1.212 1.212-6.35-3.236-11.623-8.51l2.318-2.311c0 .001 9.885 10.24 9.305 10.821z"/>`),
		g.Group(children),
	)
}