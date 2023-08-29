package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M16 6H32V12H16V6Z"/><path d="M6 42L6 6"/><path fill="#2F88FF" stroke-linejoin="round" d="M16 21H36V27H16V21Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M16 36H42V42H16V36Z"/></g>`),
		g.Group(children),
	)
}