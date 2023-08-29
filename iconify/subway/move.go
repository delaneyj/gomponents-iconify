package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 256L405.3 149.3v85.4h-128v-128h85.4L256 0L149.3 106.7h85.4v128h-128v-85.4L0 256l106.7 106.7v-85.4h128v128h-85.4L256 512l106.7-106.7h-85.4v-128h128v85.4z"/>`),
		g.Group(children),
	)
}