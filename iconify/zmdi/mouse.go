package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 7q64 8 106.5 56T341 176H192V7zM0 304v-85h341v85q0 71-50 121t-120.5 50T50 425T0 304zM149 7v169H0q0-65 43-113T149 7z"/>`),
		g.Group(children),
	)
}