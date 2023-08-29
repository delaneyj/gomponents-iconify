package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReleaseGateCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m749 403l557 557l-557 557l-90-90l402-403H0V896h1061L659 493l90-90zm787 1136V128h-384V0h512v1411l-128 128zm403-176l90 90l-557 558l-269-270l90-90l179 178l467-466z"/>`),
		g.Group(children),
	)
}