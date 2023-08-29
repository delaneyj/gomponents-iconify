package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReleaseGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m749 403l557 557l-557 557l-90-90l402-403H0V896h1061L659 493l90-90zM1152 0h512v1920h-512v-128h384V128h-384V0z"/>`),
		g.Group(children),
	)
}