package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worldexplorers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 27.026h-39a19.5 19.5 0 0 1 39 0ZM6.329 18.774h6.395m24.952-5.644H22.838m7.419-1.971h5.037M20.29 30.078h6.887m-3.444 10.396V30.078m11.88 10.396V30.078L42.5 40.474V30.078m-30.113 6.91v.042a3.444 3.444 0 0 1-3.443 3.444h0A3.444 3.444 0 0 1 5.5 37.03v-3.508a3.444 3.444 0 0 1 3.444-3.444h0a3.444 3.444 0 0 1 3.443 3.444v.042m1.776 6.879l3.45-10.365m3.307 10.396l-3.307-10.396m2.201 6.918H15.31m11.045 3.447l3.451-10.365m3.307 10.396l-3.307-10.396m2.201 6.918h-4.504M7.312 16.933h2.782"/>`),
		g.Group(children),
	)
}