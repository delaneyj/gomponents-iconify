package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 76q-84 25-171 24q-87 0-170-24v232q83-24 170-24t171 24V76zm31-55q12 0 12 14v314q0 14-12 14q-4 0-7-2q-94-35-195-35T19 361q-4 2-7 2q-12 0-12-14V35q0-14 12-14q3 0 7 2q94 35 194 35q101 0 195-35q3-2 7-2z"/>`),
		g.Group(children),
	)
}