package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortraitTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaPortrait20" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaPortrait21" stroke-opacity="1" d="m 49.616696,208.00014 c 0,0 30.383521,0 46.383451,-15.99992"/><path id="galaPortrait22" stroke-opacity="1" d="m 96.000147,192.00022 c 0,0 15.999943,15.99992 31.999853,15.99992"/><path id="galaPortrait23" stroke-opacity="1" d="m 206.38331,208.00012 c 0,0 -30.38351,2e-5 -46.38345,-15.9999"/><path id="galaPortrait24" stroke-opacity="1" d="m 159.99986,192.00022 c 0,0 -15.99994,15.99992 -31.99986,15.99992"/><path id="galaPortrait25" d="m 128,48.000858 c 31.99986,0 47.9998,15.999926 47.9998,63.999712 -1e-5,39.99983 -31.99988,63.99971 -47.9998,63.99971"/><path id="galaPortrait26" d="m 128,48.000858 c -31.999857,0 -47.999787,15.999926 -47.999783,63.999712 3e-6,39.99983 31.999873,63.99971 47.999793,63.99971"/><path id="galaPortrait27" stroke-opacity="1" d="m 159.99986,192.00022 -7.41601,-27.67694"/><path id="galaPortrait28" stroke-opacity="1" d="m 96.000147,192.00022 7.416013,-27.67694"/><circle id="galaPortrait29" cx="128" cy="128" r="111.999"/></g>`),
		g.Group(children),
	)
}