package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.832 8.556a1 1 0 1 1-1.663-1.112a1 1 0 0 1 1.663 1.112Zm.831.555A2 2 0 1 0 6.338 6.89a2 2 0 0 0 3.325 2.22Z"/>`),
		g.Group(children),
	)
}