package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 42L37 42C38.6569 42 40 40.6569 40 39L40 9C40 7.34315 38.6569 6 37 6L17 6"/><path fill="#2F88FF" d="M22 36C27.5228 36 32 31.5228 32 26H12C12 31.5228 16.4772 36 22 36Z"/><path stroke-linecap="round" d="M16 18V20"/><path stroke-linecap="round" d="M22 18V20"/><path stroke-linecap="round" d="M28 18V20"/></g>`),
		g.Group(children),
	)
}