package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserMediumTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.44 11.2a1.5 1.5 0 0 0 0 2.122l4.242 4.242a1.5 1.5 0 0 0 2.121 0l.72-.72a5.46 5.46 0 0 1-.369-1.045l-1.058 1.058a.5.5 0 0 1-.707 0l-4.243-4.242a.5.5 0 0 1 0-.707l1.69-1.69l4.165 4.164a5.5 5.5 0 0 1 7.843-4.859l.72-.72a1.5 1.5 0 0 0 0-2.121l-4.242-4.243a1.5 1.5 0 0 0-2.122 0L2.44 11.2ZM14.5 19a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/>`),
		g.Group(children),
	)
}