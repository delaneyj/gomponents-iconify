package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.133 9.16a.5.5 0 1 0 .738.675l5.631-6.168v13.831a.5.5 0 1 0 1 0V3.67l5.628 6.165a.5.5 0 0 0 .739-.674l-6.314-6.916a.746.746 0 0 0-.632-.24a.746.746 0 0 0-.476.24L3.133 9.16Z"/>`),
		g.Group(children),
	)
}