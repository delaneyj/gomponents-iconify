package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M7 9C7 7.34315 8.34315 6 10 6H41V42H10C8.34315 42 7 40.6569 7 39V9Z"/><path stroke="#fff" stroke-linecap="round" d="M7 34L41 34"/><path stroke="#000" stroke-linecap="round" d="M7 30V38"/><path stroke="#000" stroke-linecap="round" d="M41 30V38"/><path fill="#43CCF8" stroke="#fff" d="M15 6H25V26L20 22L15 26V6Z"/><path stroke="#000" stroke-linecap="round" d="M11 6L29 6"/></g>`),
		g.Group(children),
	)
}