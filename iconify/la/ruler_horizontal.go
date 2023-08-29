package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 8v16h32V8zm2 2h3v7h2v-7h2v4h2v-4h2v7h2v-7h2v4h2v-4h2v7h2v-7h2v4h2v-4h3v12H2z"/>`),
		g.Group(children),
	)
}