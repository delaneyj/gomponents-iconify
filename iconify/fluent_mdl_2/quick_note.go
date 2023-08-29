package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H640l-512-512V128h1792zM640 1739v-331H309l331 331zM1792 256H256v1024h512v512h1024V256z"/>`),
		g.Group(children),
	)
}