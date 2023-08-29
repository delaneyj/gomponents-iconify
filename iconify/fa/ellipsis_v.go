package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M384 1120v192q0 40-28 68t-68 28H96q-40 0-68-28t-28-68v-192q0-40 28-68t68-28h192q40 0 68 28t28 68zm0-512v192q0 40-28 68t-68 28H96q-40 0-68-28T0 800V608q0-40 28-68t68-28h192q40 0 68 28t28 68zm0-512v192q0 40-28 68t-68 28H96q-40 0-68-28T0 288V96q0-40 28-68T96 0h192q40 0 68 28t28 68z"/>`),
		g.Group(children),
	)
}