package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tsacdop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.46 15.682V40.04a3.46 3.46 0 0 1-6.92 0V15.682a3.46 3.46 0 1 1 6.92 0m8.534-3.768a7.34 7.34 0 1 1-5.07 13.777M12.006 18.73a7.34 7.34 0 1 1 5.07-13.777m-.001 0l18.919 6.96m-5.069 13.778L12.006 18.73"/>`),
		g.Group(children),
	)
}