package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M42 42V6"/><path fill="#2F88FF" stroke-linejoin="round" d="M16 6H32V12H16V6Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M12 21H32V27H12V21Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M6 36H32V42H6V36Z"/></g>`),
		g.Group(children),
	)
}