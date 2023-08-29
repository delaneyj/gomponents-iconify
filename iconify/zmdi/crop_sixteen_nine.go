package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSixteenNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 64q18 0 30.5 12.5T384 107v170q0 18-12.5 30.5T341 320H43q-18 0-30.5-12.5T0 277V107q0-18 12.5-30.5T43 64h298zm0 213V107H43v170h298z"/>`),
		g.Group(children),
	)
}