package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.971 11.861c.28 3.888-1.527 6.093-2.615 7.636c-1.694 1.786-3.84 4.22-9.291 4.356c-4.624.183-6.896-1.85-8.804-3.617c-2.487-2.733-3.136-4.35-3.26-8.375c-.013-2.467.939-4.929 2.602-7.095C4.934 1.474 8.64.37 12.065.143c3.592-.14 6.449 1.672 8.399 3.624c2.496 2.632 3.263 4.892 3.505 8.094Z"/>`),
		g.Group(children),
	)
}