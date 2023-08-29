package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOffTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708l15 15a.5.5 0 0 0 .708-.708l-15-15Zm13.938 9.268L15.707 12.5L7.5 4.293l1.085-1.086a2 2 0 0 1 2.829 0l5.378 5.379a2 2 0 0 1 0 2.828ZM6.792 9L11 13.207l-4.086 4.086a2 2 0 0 1-2.828 0l-1.38-1.379a2 2 0 0 1 0-2.828L6.794 9Zm.062 4.854l1-1a.5.5 0 1 0-.708-.707l-1 1a.5.5 0 1 0 .708.707Z"/>`),
		g.Group(children),
	)
}