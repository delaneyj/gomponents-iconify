package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwiftBackup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M40.444 25.179L7.556 5.5m27.792 30.21L7.556 19.08v13.58L24 42.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M40.444 25.179v7.481L24 42.5m16.444-37v7.481L30.251 19.08M7.556 5.5h32.888"/>`),
		g.Group(children),
	)
}