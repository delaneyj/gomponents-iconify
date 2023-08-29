package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h2m2-3h4.93a2 2 0 0 1 1.664.89l.812 1.22A2 2 0 0 0 18.07 7H19a2 2 0 0 1 2 2v6.5M9.5 9.877a4 4 0 1 0 5.5 5.768M3 3l18 18"/>`),
		g.Group(children),
	)
}