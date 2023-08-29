package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrapheneApps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.467 31.46V16.541a3.122 3.122 0 0 0-.091-.955a1.304 1.304 0 0 0-.203-.346a3.133 3.133 0 0 0-.788-.547L25.05 7.192a3.144 3.144 0 0 0-.875-.393a1.254 1.254 0 0 0-.349 0a3.144 3.144 0 0 0-.875.393l-13.336 7.5a3.133 3.133 0 0 0-.788.548a1.304 1.304 0 0 0-.203.346a3.122 3.122 0 0 0-.09.955V31.46a3.123 3.123 0 0 0 .09.955a1.304 1.304 0 0 0 .203.346a3.133 3.133 0 0 0 .788.547L22.95 40.81a3.144 3.144 0 0 0 .875.393a1.254 1.254 0 0 0 .349 0a3.144 3.144 0 0 0 .875-.393l13.336-7.5a3.133 3.133 0 0 0 .789-.548a1.304 1.304 0 0 0 .202-.346a3.123 3.123 0 0 0 .091-.955Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.289 15.396L24.145 24L8.71 15.396M24.145 24L24 41.207"/>`),
		g.Group(children),
	)
}