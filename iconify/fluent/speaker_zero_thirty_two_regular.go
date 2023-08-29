package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerZeroThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.866 4.72c.788-.788 2.134-.23 2.134.883v20.793c0 1.114-1.346 1.671-2.134.884l-4.694-4.695A2 2 0 0 0 9.757 22H6a4 4 0 0 1-4-4v-4a4 4 0 0 1 4-4h3.757a2 2 0 0 0 1.415-.586l4.694-4.695ZM16 7.413l-3.414 3.414a4 4 0 0 1-2.829 1.171H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h3.757a4 4 0 0 1 2.829 1.172L16 24.585V7.414Z"/>`),
		g.Group(children),
	)
}