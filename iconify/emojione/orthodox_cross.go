package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrthodoxCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#c28fef" d="M55 32v-8H36v-7h6V9h-6V2h-8v7h-6v8h6v7H9v8h19v13l-7-3v8l7 3v9h8v-6l7 3v-8l-7-3V32z"/>`),
		g.Group(children),
	)
}