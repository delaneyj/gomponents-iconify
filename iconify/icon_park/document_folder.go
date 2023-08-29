package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M32 6H22V42H32V6Z"/><path fill="#2F88FF" stroke="#000" d="M42 6H32V42H42V6Z"/><path fill="#2F88FF" stroke="#000" d="M10 6L18 7L14.5 42L6 41L10 6Z"/><path stroke="#fff" stroke-linecap="round" d="M37 18V15"/><path stroke="#fff" stroke-linecap="round" d="M27 18V15"/></g>`),
		g.Group(children),
	)
}