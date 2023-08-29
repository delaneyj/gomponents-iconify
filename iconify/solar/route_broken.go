package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouteBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g stroke="currentColor" stroke-linecap="round" stroke-width="1.5" clip-path="url(#solarRouteBroken0)"><path d="M19.071 4.929c3.333 3.333 5 5 5 7.07c0 2.072-1.667 3.739-5 7.072s-5 5-7.071 5c-2.071 0-3.738-1.667-7.071-5c-3.333-3.334-5-5-5-7.071c0-2.071 1.667-3.738 5-7.071c3.333-3.334 5-5 7.071-5c1.377 0 2.575.737 4.204 2.21"/><path stroke-linejoin="round" d="M16 11.5L13.333 9M16 11.5L13.333 14M16 11.5h-5.333C9.777 11.5 8 12 8 14"/></g><defs><clipPath id="solarRouteBroken0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}