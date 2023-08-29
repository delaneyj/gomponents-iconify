package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinkingTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 18a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 1.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM9.5 15a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm0 1.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM11.823 2a5.414 5.414 0 0 1 5.33 4.47h.082a3.765 3.765 0 1 1 0 7.53H6.412a3.765 3.765 0 1 1 0-7.53h.081A5.414 5.414 0 0 1 11.823 2Zm.006 1.498a3.927 3.927 0 0 0-3.923 3.728a.693.693 0 0 1-.692.659h-.7a2.31 2.31 0 1 0 0 4.617h10.63a2.31 2.31 0 1 0 0-4.617h-.7a.693.693 0 0 1-.692-.659a3.927 3.927 0 0 0-3.923-3.728Z"/>`),
		g.Group(children),
	)
}