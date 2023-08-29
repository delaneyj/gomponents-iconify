package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M6 14H42V24C38 32 32 36 24 36C16 36 10 32 6 24V14Z"/><path stroke="#000" stroke-linecap="round" d="M33 34L32 44H16L15 34"/><path stroke="#fff" stroke-linecap="round" d="M22 24H26"/><path stroke="#000" stroke-linecap="round" d="M16 4L16 12"/><path stroke="#000" stroke-linecap="round" d="M32 4V12"/></g>`),
		g.Group(children),
	)
}