package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphUnfold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 4V44"/><path d="M38 4V8"/><path d="M38 22V26"/><path d="M38 40V44"/><path d="M14 15H42"/><path d="M14 33H42"/></g>`),
		g.Group(children),
	)
}