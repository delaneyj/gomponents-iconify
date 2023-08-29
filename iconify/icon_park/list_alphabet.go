package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListAlphabet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 9H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 19H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 29H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 39H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 29H12L6 39H12"/><path fill="#2F88FF" d="M11 8.9999L7 9L6.3 16H11.7L11 8.9999Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 19L6.3 16M12 19L11.7 16M11.7 16L11 8.9999L7 9L6.3 16M11.7 16H6.3"/></g>`),
		g.Group(children),
	)
}