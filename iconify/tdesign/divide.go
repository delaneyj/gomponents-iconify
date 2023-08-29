package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7h2.004v2.004H11V7Zm-4.5 4h11v2h-11v-2Zm4.5 4h2.004v2.004H11V15Z"/>`),
		g.Group(children),
	)
}