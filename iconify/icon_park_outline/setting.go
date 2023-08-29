package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Setting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M36.686 15.171a15.37 15.37 0 0 1 2.529 6.102H44v5.454h-4.785a15.37 15.37 0 0 1-2.529 6.102l3.385 3.385l-3.857 3.857l-3.385-3.385a15.37 15.37 0 0 1-6.102 2.529V44h-5.454v-4.785a15.37 15.37 0 0 1-6.102-2.529l-3.385 3.385l-3.857-3.857l3.385-3.385a15.37 15.37 0 0 1-2.529-6.102H4v-5.454h4.785a15.37 15.37 0 0 1 2.529-6.102l-3.385-3.385l3.857-3.857l3.385 3.385a15.37 15.37 0 0 1 6.102-2.529V4h5.454v4.785a15.37 15.37 0 0 1 6.102 2.529l3.385-3.385l3.857 3.857l-3.385 3.385Z"/><path d="M24 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g>`),
		g.Group(children),
	)
}