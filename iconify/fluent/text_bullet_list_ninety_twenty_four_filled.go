package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.499 3.5a1.5 1.5 0 1 1-2.999 0a1.5 1.5 0 0 1 2.999 0Zm-.5 4v13.503a1 1 0 0 1-1.993.117L5 21.003V7.5a1 1 0 0 1 1.994-.116L7 7.5Zm6.5-4a1.5 1.5 0 1 1-2.999 0a1.5 1.5 0 0 1 2.999 0Zm-.5 4v13.503a1 1 0 0 1-1.993.117L11 21.003V7.5a1 1 0 0 1 1.994-.116L13 7.5Zm6.493-4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-.493 4v13.503a1 1 0 0 1-1.994.117L17 21.003V7.5a1 1 0 0 1 1.993-.116L19 7.5Z"/>`),
		g.Group(children),
	)
}