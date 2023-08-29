package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropThreeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h298zm0 299V64H43v256h298z"/>`),
		g.Group(children),
	)
}