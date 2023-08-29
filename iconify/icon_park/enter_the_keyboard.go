package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterTheKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M42 7H6C4.89543 7 4 7.89543 4 9V37C4 38.1046 4.89543 39 6 39H42C43.1046 39 44 38.1046 44 37V9C44 7.89543 43.1046 7 42 7Z"/><path stroke="#fff" stroke-linecap="round" d="M12 19H14"/><path stroke="#fff" stroke-linecap="round" d="M21 19H23"/><path stroke="#fff" stroke-linecap="round" d="M29 19H36"/><path stroke="#fff" stroke-linecap="round" d="M12 28H36"/></g>`),
		g.Group(children),
	)
}