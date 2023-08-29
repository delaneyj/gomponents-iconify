package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.357c-10.77 0-19.5 7.051-19.5 15.75a12.658 12.658 0 0 0 .087 1.396c.016.145.034.289.055.433c1.39 9.622 13.25 18.078 18.054 19.657A.998.998 0 0 0 24 41.637v-4.78c10.77 0 19.5-7.051 19.5-15.75S34.77 5.357 24 5.357Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35 14.107a7.195 7.195 0 0 0-3 0c-5.837 2.545-6.633 14.126-13 14c-3.072-.061-5.998-2.657-6-7s2.927-6.94 6-7c6.367-.127 7.163 11.454 13 14a7.193 7.193 0 0 0 3 0"/>`),
		g.Group(children),
	)
}