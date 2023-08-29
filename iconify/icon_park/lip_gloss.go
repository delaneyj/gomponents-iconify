package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LipGloss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M24 12H8C8 12 8 16 9 26C10 36 12 44 12 44H20C20 44 22 36 23 26C24 16 24 12 24 12Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M42 34H28C28 34 28 36 29 39C30 42 31.5 44 31.5 44H38.5C38.5 44 40 42 41 39C42 36 42 34 42 34Z"/><path d="M35 34V13"/><path fill="#2F88FF" stroke-linejoin="round" d="M31 7H39L37 13H33L31 7Z"/><rect width="10" height="6" x="11" y="6" fill="#2F88FF" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}