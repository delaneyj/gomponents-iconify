package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powertunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.36 36.393L19.553 22.2l8.69 8.691L43.5 14.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 17.045a22.187 22.187 0 0 1 23.588-3.51a23.235 23.235 0 0 1 13.585 20.099"/>`),
		g.Group(children),
	)
}