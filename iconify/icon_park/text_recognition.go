package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 33V39C6 40.6569 7.34315 42 9 42H15"/><path d="M33 42H39C40.6569 42 42 40.6569 42 39V33"/><path d="M42 15V9C42 7.34315 40.6569 6 39 6H33"/><path d="M6 15V9C6 7.34315 7.34315 6 9 6H15"/><path d="M24 15V35"/><path d="M17 15H24H31"/></g>`),
		g.Group(children),
	)
}