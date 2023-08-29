package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M467 435q18-19 45-19t45 19l448 451q19 19 19 46t-19 46l-26 27q-19 19-45.5 19t-45.5-19L512 626l-376 379q-19 19-45.5 19T45 1005l-26-27Q0 959 0 932t19-46zm512 155q-19 19-45.5 19T888 590L512 210L136 590q-19 19-45.5 19T45 590l-26-28Q0 543 0 516.5T19 471L467 19q18-19 45-19t45 19l448 452q19 19 19 45.5t-19 45.5z"/>`),
		g.Group(children),
	)
}