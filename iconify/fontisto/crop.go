package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.455 15.273h2.182V6.546a2.182 2.182 0 0 0-2.182-2.182H8.728v2.182h8.727zm-10.91 2.182V0H4.363v4.364H-.001v2.182h4.364v10.909c0 1.205.977 2.182 2.182 2.182h10.909v4.364h2.182v-4.364H24v-2.182z"/>`),
		g.Group(children),
	)
}