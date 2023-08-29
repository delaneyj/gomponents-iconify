package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwindonBus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 1 24 2.5A21.5 21.5 0 0 1 45.5 24ZM16.903 3.699v40.602M12.38 5.907v36.186"/><circle cx="24.923" cy="15.71" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.923 18.73v11.478m12.355 0V23.06a4.331 4.331 0 0 0-4.33-4.331h0a4.331 4.331 0 0 0-4.332 4.331v7.147"/>`),
		g.Group(children),
	)
}