package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 52H64a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h128a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12Zm4 140a4 4 0 0 1-4 4H64a4 4 0 0 1-4-4V64a4 4 0 0 1 4-4h128a4 4 0 0 1 4 4Zm40-136v144a4 4 0 0 1-8 0V56a4 4 0 0 1 8 0ZM28 56v144a4 4 0 0 1-8 0V56a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}