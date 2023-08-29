package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSortHorizontalLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 9.5H8m0 0L10.75 7M8 9.5l2.75 2.5M8 14.5h8m0 0L13.25 12M16 14.5L13.25 17"/></g>`),
		g.Group(children),
	)
}