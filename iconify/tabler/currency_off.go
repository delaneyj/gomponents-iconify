package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.531 14.524a7 7 0 0 0-9.06-9.053M7.049 7.053a7 7 0 0 0 9.903 9.896M4 4l3 3m13-3l-3 3M4 20l3-3m13 3l-3-3M3 3l18 18"/>`),
		g.Group(children),
	)
}