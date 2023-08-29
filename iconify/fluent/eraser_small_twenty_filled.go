package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserSmallTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.44 11.2a1.5 1.5 0 0 0 0 2.122l4.242 4.242a1.5 1.5 0 0 0 2.121 0l2.212-2.212a4 4 0 0 1 4.337-4.337l2.212-2.212a1.5 1.5 0 0 0 0-2.121l-4.242-4.243a1.5 1.5 0 0 0-2.122 0L2.44 11.2Zm.706 1.415a.5.5 0 0 1 0-.707l1.69-1.69l4.95 4.95l-1.69 1.69a.5.5 0 0 1-.707 0l-4.243-4.243ZM15 18a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}