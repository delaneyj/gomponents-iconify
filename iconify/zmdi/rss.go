package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M47 269q19 0 33 13.5t14 33T80 349t-33 14t-33-14t-14-33.5t14-33T47 269zM0 146q90 0 153.5 63.5T217 363h-62q0-64-45.5-109.5T0 208v-62zM0 21q93 0 171.5 46t124 124.5T341 363h-62q0-116-81.5-198T0 83V21z"/>`),
		g.Group(children),
	)
}