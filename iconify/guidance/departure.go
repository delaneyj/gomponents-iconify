package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Departure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 20.5h24M9.019 10.193L1.417 7.188C3.185 4.713 4.422 1 4.422 1l1.238.53s.246 1.343 0 3.536l12.398 3.869a4.074 4.074 0 0 1 2.64 2.604c.373 1.116.843 2.473 1.225 3.427l-.353.353l-8.09-3.011c-3.931 2.304-7.467 3.011-7.467 3.011l-1.555-1.555s2.44-1.45 4.56-3.571Z"/>`),
		g.Group(children),
	)
}