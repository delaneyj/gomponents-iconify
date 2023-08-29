package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiMonitorPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.021 32.921a4.216 4.216 0 1 0 4.216 4.217v0h0a4.216 4.216 0 0 0-4.216-4.217Zm1.271.196l8.039-25.015m-2.566 20.514c-.091-.077-.188-.142-.281-.215m-6.639-2.221a10.486 10.486 0 0 0-6.515 2.399v.036m19.848-7.289a20.654 20.654 0 0 0-3.651-2.393m-6.671-2.122a20.642 20.642 0 0 0-15.992 4.515M43.5 13.674a30.552 30.552 0 0 0-6.898-4.31M29.93 7.249A30.532 30.532 0 0 0 4.5 13.674"/>`),
		g.Group(children),
	)
}