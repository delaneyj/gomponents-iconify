package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 0v2048h-128V0h128zM512 384h1152v512H512V384zm128 384h896V512H640v256zM0 1152h1664v512H0v-512zm128 384h1408v-256H128v256z"/>`),
		g.Group(children),
	)
}