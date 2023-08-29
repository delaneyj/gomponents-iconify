package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 36v8h32V4H8v8M6 30h4m-4-6h4m-4-6h4"/><circle cx="24" cy="17" r="4"/><path d="M32 35a8 8 0 1 0-16 0"/></g>`),
		g.Group(children),
	)
}