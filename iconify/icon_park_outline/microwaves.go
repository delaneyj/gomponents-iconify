package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microwaves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="24" height="38" x="5" y="34" stroke="currentColor" stroke-width="4" rx="2" transform="rotate(-90 5 34)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 19h4m-11 6h22m-23 9v4m8-4v4m8-4v4m8-4v4"/><path fill="currentColor" d="M15 19a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g>`),
		g.Group(children),
	)
}