package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128V43h85v85H0zm0 107v-86h85v86H0zm107 0v-86h85v86h-85zm106 0v-86h86v86h-86zM107 128V43h85v85h-85zm106-85h86v85h-86V43zm107 192v-86h85v86h-85zM0 341v-85h85v85H0zm107 0v-85h85v85h-85zm106 0v-85h86v85h-86zm107 0v-85h85v85h-85zm0-298h85v85h-85V43z"/>`),
		g.Group(children),
	)
}