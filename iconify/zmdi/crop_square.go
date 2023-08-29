package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 21q17 0 29.5 12.5T341 64v256q0 18-12.5 30.5T299 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h256zm0 299V64H43v256h256z"/>`),
		g.Group(children),
	)
}