package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentificationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 8.5A2.5 2.5 0 0 1 5.5 6h12A2.5 2.5 0 0 1 20 8.5v7a2.5 2.5 0 0 1-2.5 2.5h-12A2.5 2.5 0 0 1 3 15.5v-7Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M4 5a1.5 1.5 0 0 0-1.5 1.5v7A1.5 1.5 0 0 0 4 15h12a1.5 1.5 0 0 0 1.5-1.5v-7A1.5 1.5 0 0 0 16 5H4ZM1.5 6.5A2.5 2.5 0 0 1 4 4h12a2.5 2.5 0 0 1 2.5 2.5v7A2.5 2.5 0 0 1 16 16H4a2.5 2.5 0 0 1-2.5-2.5v-7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.182 7.818a.5.5 0 0 1 .5-.5h2.727a.5.5 0 0 1 0 1H9.682a.5.5 0 0 1-.5-.5Zm0 4.092a.5.5 0 0 1 .5-.5h5.454a.5.5 0 1 1 0 1H9.682a.5.5 0 0 1-.5-.5Zm0-2.728a.5.5 0 0 1 .5-.5h5.454a.5.5 0 0 1 0 1H9.682a.5.5 0 0 1-.5-.5Zm0 1.364a.5.5 0 0 1 .5-.5h5.454a.5.5 0 1 1 0 1H9.682a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M7.773 9.182a1.364 1.364 0 1 1-2.728 0a1.364 1.364 0 0 1 2.728 0Z"/><path d="M8.045 11.775c0 .688-.732.623-1.636.623c-.904 0-1.636.066-1.636-.623c0-.688.732-1.557 1.636-1.557c.904 0 1.636.87 1.636 1.557ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}