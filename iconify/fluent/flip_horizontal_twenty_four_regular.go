package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.88 19.656a.75.75 0 0 1-.63.344h-7.5a.75.75 0 0 1-.75-.75V2.75a.75.75 0 0 1 1.433-.31l7.5 16.5a.75.75 0 0 1-.052.716ZM14.5 6.213V18.5h5.585L14.5 6.213ZM2.5 20a.5.5 0 0 1-.452-.713l8-17A.5.5 0 0 1 11 2.5v17a.5.5 0 0 1-.5.5h-8Z"/>`),
		g.Group(children),
	)
}