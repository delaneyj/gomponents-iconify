package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentificationCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7 8a1.5 1.5 0 0 0-1.5 1.5v7A1.5 1.5 0 0 0 7 18h12a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 19 8H7ZM4.5 9.5A2.5 2.5 0 0 1 7 7h12a2.5 2.5 0 0 1 2.5 2.5v7A2.5 2.5 0 0 1 19 19H7a2.5 2.5 0 0 1-2.5-2.5v-7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.182 10.818a.5.5 0 0 1 .5-.5h2.727a.5.5 0 0 1 0 1h-2.727a.5.5 0 0 1-.5-.5Zm0 4.092a.5.5 0 0 1 .5-.5h5.454a.5.5 0 1 1 0 1h-5.454a.5.5 0 0 1-.5-.5Zm0-2.728a.5.5 0 0 1 .5-.5h5.454a.5.5 0 0 1 0 1h-5.454a.5.5 0 0 1-.5-.5Zm0 1.364a.5.5 0 0 1 .5-.5h5.454a.5.5 0 1 1 0 1h-5.454a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M10.773 12.182a1.364 1.364 0 1 1-2.728 0a1.364 1.364 0 0 1 2.728 0Z"/><path d="M11.045 14.775c0 .688-.732.623-1.636.623c-.904 0-1.636.066-1.636-.623c0-.688.732-1.557 1.636-1.557c.904 0 1.636.87 1.636 1.557Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}