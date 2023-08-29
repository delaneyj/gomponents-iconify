package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KoalaBear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="26" r="15" fill="#2F88FF" stroke="#000" stroke-width="4"/><path fill="#43CCF8" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M21 31.5625C21 29.664 24 24 24 24C24 24 27 29.664 27 31.5625C27 33.461 25.6569 35 24 35C22.3431 35 21 33.461 21 31.5625Z"/><circle cx="17.039" cy="23.493" r="2" fill="#fff"/><circle cx="31.039" cy="23.493" r="2" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 11.9269C15.6214 9.52446 8.49451 6.09319 5.34959 9.30595C4.01392 10.4178 2.98594 15.4499 5.72704 23L10 19.9798"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 11.9269C32.3786 9.52446 39.5055 6.0932 42.6504 9.30595C43.9861 10.4178 45.0141 15.4499 42.273 23L38 19.9802"/></g>`),
		g.Group(children),
	)
}