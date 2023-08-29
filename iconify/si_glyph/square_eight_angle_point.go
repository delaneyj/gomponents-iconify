package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareEightAnglePoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 4V1.041h-2.959v1.021H10V1.041H6.041v1.021H3V1.041H.041V4h1.011v3.041H.041V11h1.011v3.041H.041V17H3v-1.021h3.041V17H10v-1.021h3.041V17H16v-2.959h-1.031V11H16V7.041h-1.031V4H16zM9.062 1.938v1.125H6.978V1.938h2.084zm6 0v1.125h-1.084V1.938h1.084zm-14.083 0h1.084v1.125H.979V1.938zm-.01 6h1.125v2.125H.969V7.938zm1.093 8.146H.978v-1.125h1.084v1.125zm4.917-.022v-1.125h2.084v1.125H6.979zm6.989 0v-1.114h1.126v1.114h-1.126zm1.126-5.958h-1.125V7.938h1.125v2.166zm-2.053-3.063V11h1.011v3.041h-1.011v1.021H10v-1.021H6.041v1.021H3v-1.021H1.969V11H3V7.041H1.969V4H3V2.979h3.041V4H10V2.979h3.041V4h1.011v3.041h-1.011z"/>`),
		g.Group(children),
	)
}