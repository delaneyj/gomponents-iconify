package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 8v16h22v-3.375l6 3V8.375l-6 3V8zm2 2h18v12H4zm24 1.625v8.75l-4-2v-4.75z"/>`),
		g.Group(children),
	)
}