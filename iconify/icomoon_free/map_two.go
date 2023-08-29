package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m10.5 3l-5-2L0 3v12l5.5-2l5 2l5.5-2V1l-5.5 2zM6 2.277l4 1.6v9.846l-4-1.6V2.277zM1 3.7l4-1.455v9.872l-4 1.454V3.699zm14 8.6l-4 1.455V3.883l4-1.455V12.3z"/>`),
		g.Group(children),
	)
}