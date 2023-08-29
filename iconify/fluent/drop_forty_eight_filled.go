package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 4c.331 0 .65.132.884.366l.235.233c1.135 1.125 4.249 4.21 7.173 8.228C35.42 17.124 38.5 22.71 38.5 28.25c0 5.376-1.599 9.357-4.312 11.99C31.485 42.865 27.829 44 24 44c-3.83 0-7.485-1.136-10.188-3.76C11.1 37.608 9.5 33.627 9.5 28.25c0-5.535 3.013-11.057 6.113-15.328c2.89-3.981 5.996-7.061 7.235-8.29l.268-.266A1.25 1.25 0 0 1 24 4Z"/>`),
		g.Group(children),
	)
}