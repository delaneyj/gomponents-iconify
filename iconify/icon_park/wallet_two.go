package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M4 8H44V16C44 16 34 18 34 24C34 30 44 32 44 32V40H4V8Z"/><path stroke-linecap="round" d="M44 24H42"/></g>`),
		g.Group(children),
	)
}