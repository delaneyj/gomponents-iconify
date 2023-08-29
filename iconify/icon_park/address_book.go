package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 36V44H40V4H8V12"/><path d="M6 30H10"/><path d="M6 24H10"/><path d="M6 18H10"/><circle cx="24" cy="17" r="4" fill="#2F88FF"/><path d="M32 35C32 30.5817 28.4183 27 24 27C19.5817 27 16 30.5817 16 35"/></g>`),
		g.Group(children),
	)
}