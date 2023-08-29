package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBtcThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M162.27 117.21A40 40 0 0 0 140 44V24a4 4 0 0 0-8 0v20h-24V24a4 4 0 0 0-8 0v20H64a4 4 0 0 0 0 8h12v144H64a4 4 0 0 0 0 8h36v20a4 4 0 0 0 8 0v-20h24v20a4 4 0 0 0 8 0v-20h12a44 44 0 0 0 10.27-86.79ZM84 52h56a32 32 0 0 1 0 64H84Zm68 144H84v-72h68a36 36 0 0 1 0 72Z"/>`),
		g.Group(children),
	)
}