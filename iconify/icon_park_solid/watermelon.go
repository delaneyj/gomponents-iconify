package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watermelon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWatermelon0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M23 34c11.046 0 20-8.954 20-20H3c0 11.046 8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="M23 23v3m-6.364-5.636l-2.121 2.121m14.849-2.121l2.122 2.121"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWatermelon0)"/>`),
		g.Group(children),
	)
}