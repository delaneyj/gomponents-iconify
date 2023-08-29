package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TradeMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4d5357" d="M2 2v7.5h10.3V34h7.5V9.5h10.3V2zm52.5 0l-6.6 13.2L41.4 2h-7.5v32h7.5V20.8L47.9 34l6.6-13.2V34H62V2z"/>`),
		g.Group(children),
	)
}