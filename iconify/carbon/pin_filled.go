package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.586 13.314L30 11.9L20 2l-1.314 1.415l1.186 1.186L8.38 14.322l-1.716-1.715L5.25 14l5.657 5.677L2 28.583L3.41 30l8.911-8.909L18 26.748l1.393-1.414l-1.716-1.716l9.724-11.49Z"/>`),
		g.Group(children),
	)
}