package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrademarkLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm-18-114a6 6 0 0 1-6 6H94v42a6 6 0 0 1-12 0v-42H72a6 6 0 0 1 0-12h32a6 6 0 0 1 6 6Zm80 0v48a6 6 0 0 1-12 0v-32l-17.48 20a6 6 0 0 1-9 0L134 120v32a6 6 0 0 1-12 0v-48a6 6 0 0 1 10.52-4L156 126.89l23.48-26.84A6 6 0 0 1 190 104Z"/>`),
		g.Group(children),
	)
}