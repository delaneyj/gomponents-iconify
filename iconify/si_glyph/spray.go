package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.254.021l-.223.948l1.953.015l-1.73-.963zm7.709 7.003v-.95S11.3 3.04 10.936 3.04h-.967L9.985.015H8.016v1.966h1v1.06h-.951c-.363 0-2.059 3.033-2.059 3.033v9.269c0 .37.324.672.725.672h5.509c.399 0 .724-.302.724-.672V7.045h.018l-.019-.02zm-.947 8.007h-1.022V7h1.022v8.031zm-6.989-13l.232.914l1.679-.883l-1.911-.031z"/>`),
		g.Group(children),
	)
}