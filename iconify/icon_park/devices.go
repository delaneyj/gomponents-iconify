package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M23 43H43V5H14V15"/><path fill="#2F88FF" stroke="#000" d="M5 15H23V43H5L5 15Z"/><path stroke="#fff" stroke-linecap="round" d="M13 37H15"/><path stroke="#000" stroke-linecap="round" d="M28 37H30"/></g>`),
		g.Group(children),
	)
}