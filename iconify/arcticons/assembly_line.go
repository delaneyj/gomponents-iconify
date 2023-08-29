package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssemblyLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.983 38.87v-4.924l16.745-16.744l2.624 1.205l.148 4.717l-17.033 17.032l-2.484-1.286z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.467 40.156v-5.101l-2.484-1.109m2.484 1.109l16.885-16.648"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.983 38.382L8.346 30.886v-3.637l17.787-17.033l13.395 7.186M22.983 34.568L8.346 27.249m4.669-4.472l13.894 7.244m-11.245-9.603l13.574 7.274m-8.972-11.859l13.718 7.113m-11.318-9.41l14.201 6.527"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.133 10.216v-.887l-2.949-1.485l-17.565 17.52l-.119 4.99l2.846 1.419v-.887m-2.727-5.522l3.135 1.493"/>`),
		g.Group(children),
	)
}