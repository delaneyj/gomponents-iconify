package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.908 2.082a5 5 0 0 0-5.833 5.783L2.22 12.72a.75.75 0 0 0-.22.53v4c0 .414.336.75.75.75h3.5a.75.75 0 0 0 .75-.75V15.5h2.25a.75.75 0 0 0 .75-.75v-1.69l1.458-1.457A6.473 6.473 0 0 1 10 7.5a6.494 6.494 0 0 1 2.908-5.418ZM11 7.5a5.5 5.5 0 1 1 8.5 4.61v2.58l1.276 1.28a.75.75 0 0 1 0 1.06l-.967.97l.967.97a.75.75 0 0 1-.046 1.102l-3.245 2.75a.75.75 0 0 1-.99-.017l-2.75-2.5a.75.75 0 0 1-.245-.555v-7.64A5.496 5.496 0 0 1 11 7.5Zm6.5-1a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}