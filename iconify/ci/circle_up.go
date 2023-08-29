package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10c-.006 5.52-4.48 9.994-10 10Zm-.016-2H12a8 8 0 1 0-.016 0ZM13 17h-2v-6.17l-2.59 2.58L7 12l5-5l5 5l-1.41 1.41L13 10.83V17Z"/>`),
		g.Group(children),
	)
}