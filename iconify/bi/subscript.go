package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m3.266 12.496l.96-2.853H7.76l.96 2.853H10L6.62 3H5.38L2 12.496h1.266Zm2.748-8.063l1.419 4.23h-2.88l1.426-4.23h.035Zm6.132 7.203v-.075c0-.332.234-.618.619-.618c.354 0 .618.256.618.58c0 .362-.271.649-.52.898l-1.788 1.832V15h3.59v-.958h-1.923v-.045l.973-1.04c.415-.438.867-.845.867-1.547c0-.8-.701-1.41-1.787-1.41c-1.23 0-1.795.8-1.795 1.576v.06h1.146Z"/>`),
		g.Group(children),
	)
}