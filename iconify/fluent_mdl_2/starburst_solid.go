package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarburstSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 1024l-384 256l96 480l-480-96l-256 384l-256-384l-480 96l96-480L0 1024l384-256l-96-480l480 96L1024 0l256 384l480-96l-96 480l384 256z"/>`),
		g.Group(children),
	)
}