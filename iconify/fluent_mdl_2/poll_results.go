package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollResults(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 640v640h-768v512H128v128H0V0h128v128h1408v512h512zM128 256v384h1280V256H128zm1024 1408v-384H128v384h1024zm768-512V768H128v384h1792z"/>`),
		g.Group(children),
	)
}