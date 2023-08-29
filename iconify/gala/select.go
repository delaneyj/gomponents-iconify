package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Select(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaSelect0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaSelect1" d="M 16.000736,48.000563 A 31.999783,31.999783 0 0 1 48.000519,16.000782"/><path id="galaSelect2" d="m -239.99926,48.000563 a 31.999783,31.999783 0 0 1 31.99978,-31.999781" transform="scale(-1 1)"/><path id="galaSelect3" d="m -239.99926,-207.99947 a 31.999783,31.999783 0 0 1 31.99978,-31.99978" transform="scale(-1)"/><path id="galaSelect4" d="m 16.000736,-207.99947 a 31.999783,31.999783 0 0 1 31.999783,-31.99978" transform="scale(1 -1)"/><path id="galaSelect5" stroke-opacity="1" d="m 239.99923,143.99987 v 31.9998"/><path id="galaSelect6" stroke-opacity="1" d="M 239.99923,80.000312 V 112.00011"/><path id="galaSelect7" stroke-opacity="1" d="m 16.000747,143.99991 v 31.99976"/><path id="galaSelect8" stroke-opacity="1" d="M 16.000751,80.000312 V 112.00011"/><path id="galaSelect9" stroke-opacity="1" d="M 112.00008,16.000744 H 80.000284"/><path id="galaSelecta" stroke-opacity="1" d="M 175.99964,16.000748 H 143.99988"/><path id="galaSelectb" stroke-opacity="1" d="M 112.00008,239.99922 H 80.000284"/><path id="galaSelectc" stroke-opacity="1" d="M 175.99964,239.99922 H 143.99988"/><path id="galaSelectd" stroke-opacity="1" d="M 96.000202,127.99999 H 159.99976"/><path id="galaSelecte" stroke-opacity="1" d="M 128,96.000192 V 159.99979"/></g>`),
		g.Group(children),
	)
}