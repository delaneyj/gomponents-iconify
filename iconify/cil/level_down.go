package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m171.313 348.686l-22.626 22.628L272 494.627l123.313-123.313l-22.626-22.628L288 433.373V96H72v32h184v305.373l-84.687-84.687z"/>`),
		g.Group(children),
	)
}