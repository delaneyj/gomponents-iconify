package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m257 271.5l256 128v-256l-256 128zm-256 0l256 128v-256L1 271.5z"/>`),
		g.Group(children),
	)
}