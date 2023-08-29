package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eventbrite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 7.735a8.881 8.881 0 0 1 9.068 3.333L7.469 14.667C8 11.334 10.53 8.531 14 7.735zm9.197 13.068C22 22.531 20.136 23.864 18 24.266c-3.599.803-7.068-.667-9.068-3.463l15.599-3.469l2.537-.531L32 15.735c0-1.068-.136-2.136-.401-3.068C29.599 4 21.068-1.197 12.402.803s-14 10.395-12 18.796s10.531 13.735 19.197 11.735c5.068-1.197 9.068-4.531 10.932-8.803c.136-.129-7.333-1.728-7.333-1.728z"/>`),
		g.Group(children),
	)
}