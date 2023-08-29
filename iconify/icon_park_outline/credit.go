package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Credit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 15V7H10v8m33 12V15H5v26h38"/><path d="m29 27l-5 7h19M26 21h-4m-6 0h-4"/></g>`),
		g.Group(children),
	)
}