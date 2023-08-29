package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="4" d="M11.064 10.414c5.543-.276 9.854-.414 12.934-.414s7.393.138 12.939.415a6 6 0 0 1 5.68 5.492c.254 3.034.381 5.706.381 8.017c0 2.339-.13 5.048-.39 8.128a6 6 0 0 1-5.587 5.483c-4.741.31-9.082.465-13.023.465c-3.94 0-8.28-.155-13.018-.465a6 6 0 0 1-5.587-5.48c-.263-3.103-.395-5.814-.395-8.131c0-2.29.129-4.963.385-8.02a6 6 0 0 1 5.68-5.49Z"/><path stroke-width="3.429" d="M21 19.61v8.796a.857.857 0 0 0 1.33.715l6.597-4.36a.857.857 0 0 0 .006-1.427l-6.598-4.436a.857.857 0 0 0-1.335.711Z"/></g>`),
		g.Group(children),
	)
}