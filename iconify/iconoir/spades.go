package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spades(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 14.5c3 4.5 9 4.47 9-.5c0-4-4-7-9-12c-5 5-9 8-9 12c0 4.97 6 5 9 .5Z"/><path d="m11.47 15.493l-3 5.625A.6.6 0 0 0 9 22h6a.6.6 0 0 0 .53-.882l-3-5.625a.6.6 0 0 0-1.06 0Z"/></g>`),
		g.Group(children),
	)
}