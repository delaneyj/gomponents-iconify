package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeniusHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.688 21.164a21.5 21.5 0 1 1 .06 6.086"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.31 21.168a18.789 18.789 0 1 1 .07 6.088"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.501 22.598h11.728v3.253H2.501zM13.091 11.9h5.612m-7.015 2.117h8.417m-8.885 2.118h9.352m-8.896 2.118h8.418M13.09 20.37h5.612"/>`),
		g.Group(children),
	)
}