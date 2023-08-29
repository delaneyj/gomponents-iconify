package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareBottomDownBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m11 13l-8 8m0 0h6m-6 0v-6"/><path d="M22 12c0 4.714 0 7.071-1.465 8.535C19.178 21.893 17.055 21.992 13 22M2 11c.008-4.055.107-6.178 1.464-7.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464c.974.974 1.3 2.343 1.41 4.536"/></g>`),
		g.Group(children),
	)
}