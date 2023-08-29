package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrival(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 20.5h24m-13.5-9.76L3 13.99c-.5-3-2.25-6.5-2.25-6.5L2 6.99s1.123.775 2.5 2.5l11.502-6.03a4.074 4.074 0 0 1 3.707-.025c1.054.525 2.346 1.152 3.291 1.556v.5l-7.85 3.59C14 13.49 12 16.49 12 16.49H9.8s.7-2.75.7-5.75Z"/>`),
		g.Group(children),
	)
}