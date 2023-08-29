package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.168 5.921c.009.006 1.877 1.889 1.885 1.896l1.314-1.32c3.855 4.404 2.466 7.059 2.615 6.908c1.901-1.899.89-5.99-2.256-9.136C9.58 1.123 5.488.112 3.591 2.011c-.148.146 2.5-1.271 6.908 2.593L9.168 5.921zm4.972-4.317l.88.877l-.946.947l-.879-.877z"/><path d="M8.873 7.166L.289 15.752a.667.667 0 0 0-.009.943a.668.668 0 0 0 .944-.01l8.584-8.584l-.935-.935z"/></g>`),
		g.Group(children),
	)
}