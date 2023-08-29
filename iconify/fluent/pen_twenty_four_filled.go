package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.891 3.047a3.578 3.578 0 1 1 5.061 5.06L20.061 9a3.25 3.25 0 0 1-.005 4.592l-1.784 1.783a.75.75 0 1 1-1.06-1.06l1.783-1.784A1.75 1.75 0 0 0 19 10.06l-9.998 10a3.106 3.106 0 0 1-1.477.825L2.924 21.98a.75.75 0 0 1-.904-.903l1.096-4.602c.133-.559.419-1.07.825-1.476l11.95-11.952Z"/>`),
		g.Group(children),
	)
}