package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListRtlNinetyTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 17a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm-1-3a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 1 0v11a.5.5 0 0 1-.5.5Zm5 0a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 1 0v11a.5.5 0 0 1-.5.5Zm4.5-.5a.5.5 0 0 0 1 0v-11a.5.5 0 0 0-1 0v11Zm.5 2.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm-4 1a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}