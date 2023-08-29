package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cooking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" d="M6 42L42 42"/><path stroke="#000" stroke-linecap="round" d="M6 36L42 36"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M9 25C9 16.7157 15.7157 10 24 10C32.2843 10 39 16.7157 39 25V36H9L9 25Z"/><path stroke="#fff" stroke-linecap="round" d="M17 25V29"/><path stroke="#000" d="M28 10V8C28 5.79086 26.2091 4 24 4V4C21.7909 4 20 5.79086 20 8V10"/></g>`),
		g.Group(children),
	)
}