package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinkingTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 18a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm5.5-3a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5ZM12 2a5.414 5.414 0 0 1 5.33 4.47h.082a3.765 3.765 0 1 1 0 7.53H6.588a3.765 3.765 0 1 1 0-7.53h.082A5.414 5.414 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}