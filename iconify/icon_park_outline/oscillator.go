package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oscillator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 9v15a6 6 0 0 0 6 6h22a6 6 0 0 0 6-6V9"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 10a6 6 0 0 1 6-6h22a6 6 0 0 1 0 12H13a6 6 0 0 1-6-6Z"/><circle cx="15" cy="10" r="2" fill="currentColor"/><circle cx="21" cy="10" r="2" fill="currentColor"/><circle cx="27" cy="10" r="2" fill="currentColor"/><circle cx="33" cy="10" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 30v14m10-14v14"/></g>`),
		g.Group(children),
	)
}