package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zlib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 26.715V6.044m-2.944 0l2.889 20.471m-5.834-20.77L24 26.715m1.256 0l7.521-21.169L38.206 4.5l-10.21 26.946h-7.992L9.794 4.5l5.429 1.046l7.521 21.169h2.512zm1.096 4.731V43.5l-2.297-2.297l-2.407 2.297V31.446m5.296-25.402l-2.889 20.471m5.834-20.77L24 26.715"/>`),
		g.Group(children),
	)
}