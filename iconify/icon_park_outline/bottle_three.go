package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.188 10h5.625L33 21.18V44H15V21.18L21.188 10ZM20 4h8v6h-8z"/><rect width="6" height="12" x="21" y="23" rx="3"/></g>`),
		g.Group(children),
	)
}