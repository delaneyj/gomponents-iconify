package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyCircleDollarLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm38-70a26 26 0 0 1-26 26h-6v10a6 6 0 0 1-12 0v-10h-18a6 6 0 0 1 0-12h36a14 14 0 0 0 0-28h-24a26 26 0 0 1 0-52h6V72a6 6 0 0 1 12 0v10h18a6 6 0 0 1 0 12h-36a14 14 0 0 0 0 28h24a26 26 0 0 1 26 26Z"/>`),
		g.Group(children),
	)
}