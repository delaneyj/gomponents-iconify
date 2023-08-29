package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" stroke-linejoin="round" d="M6 15H42L40 42H8L6 15Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 19V6H32V19"/><path stroke="#fff" stroke-linecap="round" d="M16 34H32"/></g>`),
		g.Group(children),
	)
}