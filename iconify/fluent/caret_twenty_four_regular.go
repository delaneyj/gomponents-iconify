package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M16.5 7.81L7.81 16.5h8.44a.25.25 0 0 0 .25-.25V7.81zm-.634-1.487c.788-.787 2.134-.23 2.134.884v9.043A1.75 1.75 0 0 1 16.25 18H7.207c-1.114 0-1.671-1.346-.884-2.134l9.543-9.543z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}