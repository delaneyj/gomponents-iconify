package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M13 38H41V16H30V4H13V38Z"/><path stroke="#000" stroke-linejoin="round" d="M30 4L41 16"/><path stroke="#000" stroke-linejoin="round" d="M7 20V44H28"/><path stroke="#fff" d="M19 20H23"/><path stroke="#fff" d="M19 28H31"/></g>`),
		g.Group(children),
	)
}