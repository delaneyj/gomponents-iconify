package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LatticePattern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 8H10"/><path d="M14 16H18"/><path d="M16 6V10"/><path d="M8 14V18"/><path d="M22 8H26"/><path d="M30 16H34"/><path d="M32 6V10"/><path d="M24 14V18"/><path d="M38 8H42"/><path d="M40 14V18"/><path d="M6 24H10"/><path d="M14 32H18"/><path d="M16 22V26"/><path d="M8 30V34"/><path d="M22 24H26"/><path d="M30 32H34"/><path d="M32 22V26"/><path d="M24 30V34"/><path d="M38 24H42"/><path d="M40 30V34"/><path d="M6 40H10"/><path d="M16 38V42"/><path d="M22 40H26"/><path d="M32 38V42"/><path d="M38 40H42"/></g>`),
		g.Group(children),
	)
}