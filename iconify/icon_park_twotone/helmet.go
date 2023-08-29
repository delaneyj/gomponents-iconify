package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Helmet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHelmet0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M42.34 32c1.07-2.45 1.66-5.16 1.66-8c0-11.05-8.95-20-20-20S4 12.95 4 24h20l.01 8c0 3.87 3.13 7 7 7s7-3.13 7-7h4.33Z"/><path fill="#fff" d="M31 34c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M24 32L6 36c2.23 3.17 4.45 6.37 8 8l12-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHelmet0)"/>`),
		g.Group(children),
	)
}