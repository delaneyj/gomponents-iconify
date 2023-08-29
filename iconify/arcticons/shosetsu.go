package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shosetsu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.755 11.5v25m-9.064-19.323h18.128v7.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.849 24.697H36.15c0 5.478-1.494 8.566-3.785 8.566h-3.934m2.74-20.418a17.76 17.76 0 0 1 4.88 3.585"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}