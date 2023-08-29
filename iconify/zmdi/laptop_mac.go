package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopMac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 320h85q0 18-12.5 30.5T469 363H43q-18 0-30.5-12.5T0 320h85q-17 0-29.5-12.5T43 277V43q0-18 12.5-30.5T85 0h342q17 0 29.5 12.5T469 43v234q0 18-12.5 30.5T427 320zM85 43v234h342V43H85zm171 298q9 0 15-6t6-15t-6-15t-15-6t-15 6t-6 15t6 15t15 6z"/>`),
		g.Group(children),
	)
}