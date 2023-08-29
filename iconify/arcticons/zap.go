package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.303 34.086a9.197 9.197 0 1 0-2.735-17.982a13.85 13.85 0 1 0-13.22 17.982"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.5 26.561l-4.285 7.525h7.787l-4.286 7.525"/>`),
		g.Group(children),
	)
}