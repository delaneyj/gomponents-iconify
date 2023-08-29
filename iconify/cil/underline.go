package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M296 152h32v88a63.966 63.966 0 0 1-88 59.313V152h24v-32H136v32h32v88a96 96 0 0 0 192 0v-88h32v-32h-96ZM136 368h256v32H136z"/>`),
		g.Group(children),
	)
}