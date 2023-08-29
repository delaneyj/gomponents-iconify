package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoMpThreeConverter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.396 9.589c-8.374-6.71-20.602-5.362-27.312 3.012c-6.71 8.374-5.362 20.602 3.012 27.313c7.729 6.193 18.879 5.582 25.884-1.418"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="33.996" cy="29.488" r="4.512"/><path d="M38.508 29.488V14h4.008"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.676 3.841l1.72 5.748l-5.748 1.721"/>`),
		g.Group(children),
	)
}