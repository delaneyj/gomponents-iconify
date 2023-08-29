package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.8 9L10 3H9L7.158 9.64L5.99 4.69h-.97L3.85 9H1v.99h3.23l.49-.37l.74-2.7L6.59 12h1.03l1.87-7.04l1.46 4.68l.48.36H15V9h-3.2z"/>`),
		g.Group(children),
	)
}