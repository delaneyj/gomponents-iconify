package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkingTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 512h-120l-449 896h-72q-1 28-3 70t-8 93t-15 105t-23 108t-33 101t-44 84t-57 57t-72 22q-53 0-92-37t-67-95t-46-131t-29-144t-15-133t-6-100h-72L248 512H128V0h128v384h1536V0h128v512zm-896 1408q23-12 41-47t32-85t23-105t16-109t10-97t5-69H897q1 25 4 69t10 97t16 109t24 105t32 84t41 48zm632-1408H392l384 768h496l384-768z"/>`),
		g.Group(children),
	)
}