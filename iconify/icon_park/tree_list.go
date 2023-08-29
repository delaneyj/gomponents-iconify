package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M38 20H18V28H38V20Z"/><path fill="#2F88FF" d="M32 6H18V14H32V6Z"/><path fill="#2F88FF" d="M44 34H18V42H44V34Z"/><path stroke-linecap="round" d="M17 10H5"/><path stroke-linecap="round" d="M17 24H5"/><path stroke-linecap="round" d="M17 38H5"/><path stroke-linecap="round" d="M5 44V4"/></g>`),
		g.Group(children),
	)
}