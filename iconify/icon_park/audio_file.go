package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 44V4H31L40 14.5V44H8Z"/><path stroke="#fff" d="M32 14L26 16.9688V31.5"/><circle cx="20.5" cy="31.5" r="5.5" fill="#43CCF8" stroke="#fff"/></g>`),
		g.Group(children),
	)
}