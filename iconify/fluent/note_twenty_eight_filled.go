package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.75A3.75 3.75 0 0 1 6.75 3h14.5A3.75 3.75 0 0 1 25 6.75V15h-6.25A3.75 3.75 0 0 0 15 18.75V25H6.75A3.75 3.75 0 0 1 3 21.25V6.75ZM16.5 25h.06L25 16.56v-.06h-6.25a2.25 2.25 0 0 0-2.25 2.25V25Z"/>`),
		g.Group(children),
	)
}