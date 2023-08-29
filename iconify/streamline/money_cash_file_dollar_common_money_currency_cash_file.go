package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashFileDollarCommonMoneyCurrencyCashFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5.5h-7a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-7Z"/><path d="M8.5 5V.5l5 5H9a.5.5 0 0 1-.5-.5Zm-4-.5V3M3 8.5c0 .75.67 1 1.5 1s1.5 0 1.5-1C6 7 3 7 3 5.5c0-1 .67-1 1.5-1s1.5.38 1.5 1m-1.5 4V11m4-1.5h3"/></g>`),
		g.Group(children),
	)
}