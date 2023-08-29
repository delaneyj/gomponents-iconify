package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="4"/><path d="m15 9l4-4M5 19l4-4m0-6L5 5m14 14l-4-4"/><path stroke-linecap="round" d="M9.412 2.339a9.954 9.954 0 0 1 5.176.002c5.335 1.43 8.5 6.913 7.071 12.247c-1.43 5.335-6.912 8.5-12.247 7.071c-5.335-1.43-8.5-6.912-7.071-12.247a9.954 9.954 0 0 1 2.586-4.484"/></g>`),
		g.Group(children),
	)
}