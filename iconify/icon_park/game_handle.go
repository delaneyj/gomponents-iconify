package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameHandle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M32 18H16C9.37258 18 4 23.3726 4 30C4 36.6274 9.37258 42 16 42H32C38.6274 42 44 36.6274 44 30C44 23.3726 38.6274 18 32 18Z"/><path stroke="#fff" stroke-linecap="round" d="M16 26V34"/><path stroke="#fff" stroke-linecap="round" d="M12 30H20"/><path stroke="#000" stroke-linecap="round" d="M24 16V9.71429H32V4"/><path fill="#43CCF8" stroke="#fff" d="M32 34C34.2091 34 36 32.2091 36 30C36 27.7909 34.2091 26 32 26C29.7909 26 28 27.7909 28 30C28 32.2091 29.7909 34 32 34Z"/></g>`),
		g.Group(children),
	)
}