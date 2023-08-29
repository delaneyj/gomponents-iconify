package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><ellipse cx="11.3" rx="8.28" ry="3.17" transform="rotate(45)"/><ellipse cy="11.3" rx="8.28" ry="3.17" transform="rotate(315)"/><path d="M8 8v0"/></g>`),
		g.Group(children),
	)
}