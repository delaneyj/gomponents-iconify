package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBtcBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M177.08 114.46A48 48 0 0 0 152 37.52V24a12 12 0 0 0-24 0v12h-16V24a12 12 0 0 0-24 0v12H64a12 12 0 0 0 0 24h4v128h-4a12 12 0 0 0 0 24h24v12a12 12 0 0 0 24 0v-12h16v12a12 12 0 0 0 24 0v-12a52 52 0 0 0 25.08-97.54ZM164 84a24 24 0 0 1-24 24H92V60h48a24 24 0 0 1 24 24Zm-12 104H92v-56h60a28 28 0 0 1 0 56Z"/>`),
		g.Group(children),
	)
}