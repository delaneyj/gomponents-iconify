package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 7)"><circle cx="3.5" cy="3.5" r="3"/><path d="M6 1.5h6.5c.828 0 2 .325 2 2s-1.172 2-2 2H6"/></g>`),
		g.Group(children),
	)
}