package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassLineBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 2C7.867 2 5.8 2 5.198 3.3a2.521 2.521 0 0 0-.13.346c-.41 1.387 1.052 2.995 3.974 6.21L11 12h2l1.958-2.143c2.922-3.216 4.383-4.824 3.974-6.21a2.507 2.507 0 0 0-.13-.348C18.2 2 16.133 2 12 2Zm-2 2.75a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Z" clip-rule="evenodd"/><path d="M5.198 20.7C5.8 22 7.867 22 12 22c4.133 0 6.2 0 6.802-1.3a2.51 2.51 0 0 0 .13-.346c.41-1.387-1.052-2.995-3.974-6.21L13 12h-2l-1.958 2.143c-2.922 3.216-4.383 4.824-3.974 6.21c.035.12.078.236.13.348Z" opacity=".5"/><path d="M10 17.75a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Z"/></g>`),
		g.Group(children),
	)
}