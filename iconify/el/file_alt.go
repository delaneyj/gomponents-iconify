package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM496.729 288.794h39.697v151.538H379.468v-38.306l117.261-113.232zm92.944 0h230.859v622.412H379.468V491.748h210.205V288.794z"/>`),
		g.Group(children),
	)
}