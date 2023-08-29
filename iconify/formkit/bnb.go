package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bnb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.67 7.04L8 4.71l2.33 2.33l1.36-1.36L8 1.99L4.31 5.68l1.36 1.36ZM2 8l1.36-1.36L4.72 8L3.36 9.36L2 8Zm6 3.29L5.67 8.96l-1.36 1.35L8 14l3.69-3.69l-1.36-1.36L8 11.28ZM11.29 8l1.36-1.36L14.01 8l-1.36 1.36L11.29 8ZM8 6.62L9.38 8L8 9.38L6.62 8.01l.24-.25l.12-.12L8 6.62Z"/>`),
		g.Group(children),
	)
}