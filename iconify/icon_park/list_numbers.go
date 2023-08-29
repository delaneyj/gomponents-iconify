package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 9H42"/><path d="M20 19H42"/><path d="M20 29H42"/><path d="M20 39H42"/><path d="M6 29H12V32L6 38V39H12"/><path d="M7 11L9 9V19M9 19H7M9 19H11"/></g>`),
		g.Group(children),
	)
}