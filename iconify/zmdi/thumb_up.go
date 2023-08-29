package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 432V176h85v256H0zm469-235v43q0 8-3 16l-64 150q-11 26-39 26H171q-18 0-30.5-12.5T128 389V176q0-18 13-30L281 5l23 23q9 9 9 22l-1 7l-20 98h135q17 0 29.5 12.5T469 197z"/>`),
		g.Group(children),
	)
}