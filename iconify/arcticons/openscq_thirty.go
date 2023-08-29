package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenscqThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M38.68 30c1.58-5.97 1.82-13.8-1.92-19.14c-2.77-3.95-7.55-6.24-12.57-6.35c-5.07-.11-9.75 2.3-12.7 6.05c-4.1 5.2-3.75 13.49-2.18 19.44"/><path d="m34.08 40.44l1.98.44c1.06.24 2.11-.44 2.34-1.51l1.38-6.62c.22-1.05-.44-2.07-1.49-2.31l-2.03-.45l-2.18 10.44Zm-5.59.72c-.21 1 .8 2.01 2.24 2.28c1.45.26 2.79-.33 3-1.33l2.98-14.28c.21-1-.8-2.02-2.24-2.28c-1.45-.26-2.79.33-2.99 1.33L28.5 41.16Zm-14.58-.72l-1.98.44c-1.06.24-2.11-.44-2.34-1.51l-1.38-6.62c-.22-1.05.44-2.07 1.49-2.31l2.03-.45l2.18 10.44Z"/><path d="M19.5 41.16c.21 1-.8 2.01-2.24 2.28c-1.45.26-2.79-.33-3-1.33l-2.98-14.28c-.21-1 .8-2.02 2.24-2.28c1.45-.26 2.79.33 3 1.33l2.98 14.28Z"/></g>`),
		g.Group(children),
	)
}