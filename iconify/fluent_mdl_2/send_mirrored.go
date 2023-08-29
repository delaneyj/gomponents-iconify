package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2030 77l-220 883l220 883L51 960L2030 77zm-337 822l149-598L503 899h1190zm149 720l-147-592H515l1327 592z"/>`),
		g.Group(children),
	)
}