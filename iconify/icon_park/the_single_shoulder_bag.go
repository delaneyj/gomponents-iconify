package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheSingleShoulderBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M28 27C28 18.1813 26.7806 4 24 4C20.5714 4 20 18.1813 20 27"/><rect width="18" height="17" x="15" y="27"/><path fill="#2F88FF" d="M15 27H33L27.7059 36H19.7647L15 27Z"/></g>`),
		g.Group(children),
	)
}