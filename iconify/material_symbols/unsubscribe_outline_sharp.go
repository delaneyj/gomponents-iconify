package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsubscribeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13L4 8v10h8q0 .525.075 1.012T12.3 20H2V4h20v7.7q-.45-.225-.963-.375T20 11.1V8l-8 5Zm0-2l8-5H4l8 5Zm7 12q-2.075 0-3.538-1.463T14 18q0-2.075 1.463-3.538T19 13q2.075 0 3.538 1.463T24 18q0 2.075-1.463 3.538T19 23Zm-3-4.5h6v-1h-6v1ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}