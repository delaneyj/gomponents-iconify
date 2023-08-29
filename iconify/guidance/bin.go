package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 9v9m5-9v9m-6-13.5h-6v.25l.24 1.05A70 70 0 0 1 4.5 21.398V22.5h15v-1.102c0-5.249.59-10.48 1.76-15.598l.24-1.05V4.5h-6m-7 0V4a3.5 3.5 0 1 1 7 0v.5m-7 0h7"/>`),
		g.Group(children),
	)
}