package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDraw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M17 12L4 6V36L17 42L31 36L44 42V12L31 6L17 12Z"/><path stroke="#fff" d="M31 6V36"/><path stroke="#fff" d="M17 12V42"/><path stroke="#000" d="M10.5 9L17 12L31 6L37.5 9"/><path stroke="#000" d="M10.5 39L17 42L31 36L37.5 39"/></g>`),
		g.Group(children),
	)
}