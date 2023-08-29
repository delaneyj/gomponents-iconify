package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRice0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M24 38c9.389 0 17-7.059 17-17H7c0 9.941 7.611 17 17 17Z"/><path d="M30 21c0-5.523-4.253-10-9.5-10S11 15.477 11 21"/><path d="M39 21c0-3.314-2.766-6-6.178-6c-1.443 0-2.77.48-3.822 1.286"/><path stroke-linecap="round" d="m33 15l3-10m2 13l4-7"/><path stroke-linecap="round" stroke-linejoin="round" d="m18 37l-2 6h16l-2-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRice0)"/>`),
		g.Group(children),
	)
}