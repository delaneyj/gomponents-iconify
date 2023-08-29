package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafOneTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.037 3.048a1.75 1.75 0 0 1 2.2.224l3.712 3.712A6.999 6.999 0 0 1 12.75 18.89v2.359a.75.75 0 1 1-1.5 0V18.89A6.999 6.999 0 0 1 7.051 6.984l3.712-3.712a1.76 1.76 0 0 1 .274-.224ZM12 11a.75.75 0 0 0-.75.75v5.63a5.524 5.524 0 0 0 1.5 0v-5.63A.75.75 0 0 0 12 11Z"/>`),
		g.Group(children),
	)
}