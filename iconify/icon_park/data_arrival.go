package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataArrival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M6 42H42V6H32H30C28.6758 9.15854 26.6758 10.7378 24 10.7378C21.3242 10.7378 19.3242 9.15854 18 6H16H6V42Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M15 24L21 30L33 18"/></g>`),
		g.Group(children),
	)
}