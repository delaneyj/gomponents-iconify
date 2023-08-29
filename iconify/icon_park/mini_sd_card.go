package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniSdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linejoin="round" d="M13.9979 18.7386L8 21.9228V44H40V4H13.9979V18.7386Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M21 12V18"/><path stroke="#fff" stroke-linecap="round" d="M33 12V18"/><path stroke="#fff" stroke-linecap="round" d="M27 12V18"/></g>`),
		g.Group(children),
	)
}