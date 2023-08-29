package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.875.074H3.639L.965 12.917h3.31l.942-3.019h4.658c2.272 0 4.116-2.035 4.116-4.543v-.907c0-2.509-1.844-4.374-4.116-4.374zm.25 5.138c0 1.033-.87 1.87-1.373 1.87l-2.928.012l.832-4.129h2.096c.503 0 1.373.838 1.373 1.871v.376z"/><path d="M15.113 3.868c0 .77-.052 1.471-.052 2.11c0 3.211-1.308 5.104-4.222 5.104H6.557l-1.571 4.907h2.639l.943-3.021h4.281c2.273 0 4.115-1.721 4.115-4.213v-.902c.001-1.727-.549-3.226-1.851-3.985z"/></g>`),
		g.Group(children),
	)
}