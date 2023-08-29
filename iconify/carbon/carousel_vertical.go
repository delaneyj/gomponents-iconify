package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 10v12a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2zM8 22h16V10H8zm16 6v4h-2v-4H10v4H8v-4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2zm0-28v4a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V0h2v4h12V0z"/>`),
		g.Group(children),
	)
}