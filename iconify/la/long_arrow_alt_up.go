package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowAltUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4.094l-6.719 6.718l1.438 1.407L15 7.938V28h2V7.937l4.281 4.282l1.438-1.406z"/>`),
		g.Group(children),
	)
}