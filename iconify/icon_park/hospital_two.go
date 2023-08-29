package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M32 11H40L44 21H4L8 11H16"/><path fill="#2F88FF" stroke="#000" d="M8 21H40V43H8V21Z"/><rect width="16" height="16" x="16" y="5" fill="#2F88FF" stroke="#000"/><rect width="8" height="14" x="16" y="29" fill="#2F88FF" stroke="#fff"/><rect width="8" height="14" x="24" y="29" fill="#2F88FF" stroke="#fff"/><path stroke="#fff" d="M21 13H27"/><path stroke="#000" d="M36 43H12"/><path stroke="#fff" d="M24 16L24 10"/></g>`),
		g.Group(children),
	)
}