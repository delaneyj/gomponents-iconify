package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBannerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m231.22 59.44l-80 168a8 8 0 1 1-14.44-6.88L165.62 160H32a8 8 0 0 1-5.88-13.43l42.56-46.1L26.59 61.9A8 8 0 0 1 32 48h192a8 8 0 0 1 7.22 11.44Z"/>`),
		g.Group(children),
	)
}