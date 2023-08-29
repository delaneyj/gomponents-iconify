package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ellipsis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M371 229q0 29-20 49t-50 20t-49-20t-20-49t20-50t49-20t50 20t20 50zm-232 0q0 29-20 49t-50 20t-49-20t-20-49t20-50t49-20t50 20t20 50zm463 0q0 29-20 49t-49 20t-49-20t-21-49t21-50t49-20t49 20t20 50z"/>`),
		g.Group(children),
	)
}