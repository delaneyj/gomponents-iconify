package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvAltPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 88q17 0 29.5 12.5T469 131v256q0 17-12.5 29.5T427 429H43q-18 0-30.5-12.5T0 387V131q0-18 12.5-30.5T43 88h162l-71-70l15-15l86 85l85-85l15 15l-70 70h162zm0 299V131H43v256h384zM171 173l149 86l-149 85V173z"/>`),
		g.Group(children),
	)
}