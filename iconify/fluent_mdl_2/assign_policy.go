package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignPolicy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 1024v768h128v128H256V0h1536v1024H384zm0-896v768h1280V128H384zm768 384V384h128v256H768V384h128v128h256zm-512 640h1408v896H640v-896zm1211 128H837l507 262l507-262zM768 1920h1152v-539l-576 297l-576-297v539z"/>`),
		g.Group(children),
	)
}