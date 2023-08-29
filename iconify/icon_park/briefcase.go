package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 16C32 9.92487 28.4183 4 24 4C19.5817 4 16 9.92487 16 16"/><path fill="#2F88FF" d="M9 16H39L40 28H27V25H21V28H8L9 16Z"/><path d="M8 28L6 42H42L40 28"/><path d="M21 25H27V31H21V25Z"/></g>`),
		g.Group(children),
	)
}