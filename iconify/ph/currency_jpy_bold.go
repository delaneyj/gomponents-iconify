package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyJpyBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m209.29 55.6l-56 68.4H176a12 12 0 0 1 0 24h-36v16h36a12 12 0 0 1 0 24h-36v28a12 12 0 0 1-24 0v-28H80a12 12 0 0 1 0-24h36v-16H80a12 12 0 0 1 0-24h22.68l-56-68.4a12 12 0 1 1 18.61-15.2L128 117.05l62.71-76.65a12 12 0 1 1 18.58 15.2Z"/>`),
		g.Group(children),
	)
}