package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M39 34V16C39 14.8954 38.1046 14 37 14H34.3125H26.3438H11C9.89543 14 9 14.8954 9 16V34C9 35.1046 9.89543 36 11 36H37C38.1046 36 39 35.1046 39 34Z"/><path stroke="#000" d="M6 36H42"/><path stroke="#000" d="M6 42H42"/><path stroke="#fff" d="M9 21H15V26H21V21H27V29H33V21H39"/><path stroke="#000" d="M9 23V19"/><path stroke="#000" d="M39 23V19"/><path stroke="#000" d="M33 14V8"/><path stroke="#000" d="M24 14L24 4"/><path stroke="#000" d="M15 14L15 8"/></g>`),
		g.Group(children),
	)
}