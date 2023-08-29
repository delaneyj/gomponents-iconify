package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><ellipse cx="24" cy="20" fill="#2F88FF" stroke="#000" rx="15" ry="16"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24 14L24 36"/><path stroke="#000" d="M30 34.6689C28.1626 35.5253 26.1333 36.0003 24 36.0003C21.8667 36.0003 19.8374 35.5253 18 34.6689"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 36L24 44"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24 22L30 16"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24 29L18 23"/></g>`),
		g.Group(children),
	)
}