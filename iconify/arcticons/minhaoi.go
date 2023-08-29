package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minhaoi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.648 30.403a3.397 3.397 0 0 1-3.454-3.355v-2.171a3.456 3.456 0 0 1 6.908 0v2.178a3.414 3.414 0 0 1-3.454 3.352Z"/><circle cx="29.05" cy="18.81" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.05 21.521v8.882M5.364 22.277c.046-5.664 3.074-10.126 6.909-13.335s8.666-4.484 14.31-4.44s14.17 4.577 15.79 9.868s-4.715 8.902-5.922 13.323s.786 9.987-1.48 12.83s-5.424 3.056-8.388 2.96c-13.823-.445-21.246-9.38-21.218-21.206Z"/>`),
		g.Group(children),
	)
}