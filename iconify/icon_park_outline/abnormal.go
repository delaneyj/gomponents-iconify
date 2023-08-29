package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abnormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 16.398V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h10"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M16 14h13m-13 7h5"/><circle cx="34" cy="34" r="10" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" transform="rotate(90 34 34)"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M34 36v3"/><circle cx="34" cy="30" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}