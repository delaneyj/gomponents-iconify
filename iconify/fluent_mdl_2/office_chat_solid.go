package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeChatSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 128v512H640v859l-384 384v-475H0V128h1792zM768 768h1280v896h-256v347l-347-347H768V768z"/>`),
		g.Group(children),
	)
}