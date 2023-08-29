package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M42 41V7H6V41H42Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 7V41M42 7H4H6V41M42 7H44M42 41H44M42 41H6M6 41H4"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 7V41"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 20V28"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 23L33 20L36 23"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 28V20"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 25L15 28L18 25"/></g>`),
		g.Group(children),
	)
}