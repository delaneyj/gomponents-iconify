package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwimmingEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.004.494l-.7.502l-2.7 1.602c-.3.099-.4.598-.202.898l.6 1.002l-2.5 2l1 1.002l2-1.002L7.504 7.5l1-1.002l-3.002-3.502l3.002-1.5V.494h-.5zm1.002 2.502A1 1 0 1 0 9.003 5a1 1 0 0 0 .003-2.003zM2 7.998L0 9v1l2-1l1.5 1l2-1l2 1L9 9l2 1V9L9 7.998L7.5 9l-2-1.002L3.5 9L2 7.998z" fill="currentColor"/>`),
		g.Group(children),
	)
}