package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" d="M8 24.5957C8 25.3713 8.62871 26 9.40426 26H38.5957C39.3713 26 40 25.3713 40 24.5957V20C40 11.1634 32.8366 4 24 4C15.1634 4 8 11.1634 8 20V24.5957Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 42L24 26"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 32L15 26"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 42H15"/></g>`),
		g.Group(children),
	)
}