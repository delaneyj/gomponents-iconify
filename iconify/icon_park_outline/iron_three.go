package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 40h40l-2-16H20c-8.837 0-16 7.163-16 16ZM16 8h24l2 16m-25 8h2m6 0h2"/>`),
		g.Group(children),
	)
}