package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSevenFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 85q18 0 30.5 12.5T384 128v128q0 18-12.5 30.5T341 299H43q-18 0-30.5-12.5T0 256V128q0-18 12.5-30.5T43 85h298zm0 171V128H43v128h298z"/>`),
		g.Group(children),
	)
}