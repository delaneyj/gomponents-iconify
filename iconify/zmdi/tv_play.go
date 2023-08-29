package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 0q17 0 29.5 12.5T469 43v256q0 17-12.5 29.5T427 341H320v43H149v-43H43q-18 0-30.5-12.5T0 299V43q0-18 12.5-30.5T43 0h384zm0 299V43H43v256h384zM320 171l-149 85V85z"/>`),
		g.Group(children),
	)
}