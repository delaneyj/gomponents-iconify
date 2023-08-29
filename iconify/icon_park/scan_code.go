package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 15V6H15"/><path d="M15 42H6V33"/><path d="M42 33V42H33"/><path d="M33 6H42V15"/><path d="M10 24H38"/></g>`),
		g.Group(children),
	)
}